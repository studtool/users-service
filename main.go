package main

import (
	"os"
	"os/signal"

	"go.uber.org/dig"

	"github.com/studtool/common/utils"

	"github.com/studtool/users-service/api"
	"github.com/studtool/users-service/beans"
	"github.com/studtool/users-service/config"
	"github.com/studtool/users-service/messages"
	"github.com/studtool/users-service/repositories"
	"github.com/studtool/users-service/repositories/mongo"
)

func main() {
	c := dig.New()

	if config.RepositoriesEnabled.Value() {
		utils.AssertOk(c.Provide(mongo.NewConnection))
		utils.AssertOk(c.Invoke(func(conn *mongo.Connection) {
			if err := conn.Open(); err != nil {
				beans.Logger.Fatal(err.Error())
			} else {
				beans.Logger.Info("storage: connection open")
			}
		}))
		defer func() {
			utils.AssertOk(c.Invoke(func(conn *mongo.Connection) {
				if err := conn.Close(); err != nil {
					beans.Logger.Fatal(err)
				} else {
					beans.Logger.Info("storage: connection closed")
				}
			}))
		}()

		utils.AssertOk(c.Provide(
			mongo.NewUsersRepository,
			dig.As(new(repositories.UsersRepository)),
		))
	} else {
		utils.AssertOk(c.Provide(
			func() repositories.UsersRepository {
				return nil
			},
		))
	}

	if config.QueuesEnabled.Value() {
		utils.AssertOk(c.Provide(messages.NewQueueClient))
		utils.AssertOk(c.Invoke(func(q *messages.QueueClient) {
			if err := q.OpenConnection(); err != nil {
				beans.Logger.Fatal(err)
			} else {
				beans.Logger.Info("queue: connection open")
			}
			if err := q.Run(); err != nil {
				beans.Logger.Fatal(err)
			} else {
				beans.Logger.Info("queue: ready to receive messages")
			}
		}))
		defer func() {
			utils.AssertOk(c.Invoke(func(q *messages.QueueClient) {
				if err := q.CloseConnection(); err != nil {
					beans.Logger.Fatal(err)
				} else {
					beans.Logger.Info("queue: connection closed")
				}
			}))
		}()
	}

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Kill)
	signal.Notify(ch, os.Interrupt)

	utils.AssertOk(c.Provide(api.NewServer))
	utils.AssertOk(c.Invoke(func(srv *api.Server) {
		go func() {
			if err := srv.Run(); err != nil {
				beans.Logger.Fatal(err)
				ch <- os.Interrupt
			}
		}()
	}))
	defer func() {
		utils.AssertOk(c.Invoke(func(srv *api.Server) {
			if err := srv.Shutdown(); err != nil {
				beans.Logger.Fatal(err)
			}
		}))
	}()

	<-ch
}
