package main

import (
	"github.com/studtool/users-service/mq"
	"os"
	"os/signal"

	"go.uber.org/dig"

	"github.com/studtool/users-service/api"
	"github.com/studtool/users-service/beans"
	"github.com/studtool/users-service/config"
	"github.com/studtool/users-service/repositories"
	"github.com/studtool/users-service/repositories/mongo"
)

func main() {
	c := dig.New()

	if config.RepositoriesEnabled.Value() {
		panicOnErr(c.Provide(mongo.NewConnection))
		panicOnErr(c.Invoke(func(conn *mongo.Connection) {
			if err := conn.Open(); err != nil {
				beans.Logger.Fatal(err.Error())
			} else {
				beans.Logger.Info("storage: connection open")
			}
		}))
		defer func() {
			panicOnErr(c.Invoke(func(conn *mongo.Connection) {
				if err := conn.Close(); err != nil {
					beans.Logger.Fatal(err)
				} else {
					beans.Logger.Info("storage: connection closed")
				}
			}))
		}()

		panicOnErr(c.Provide(
			mongo.NewUsersRepository,
			dig.As(new(repositories.UsersRepository)),
		))
	} else {
		panicOnErr(c.Provide(
			func() repositories.UsersRepository {
				return nil
			},
		))
	}

	if config.QueuesEnabled.Value() {
		panicOnErr(c.Provide(mq.NewQueue))
		panicOnErr(c.Invoke(func(q *mq.MQ) {
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
			panicOnErr(c.Invoke(func(q *mq.MQ) {
				if err := q.CloseConnection(); err != nil {
					beans.Logger.Fatal(err)
				} else {
					beans.Logger.Info("queue: connection closed")
				}
			}))
		}()
	}

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)

	panicOnErr(c.Provide(api.NewServer))
	panicOnErr(c.Invoke(func(srv *api.Server) {
		go func() {
			if err := srv.Run(); err != nil {
				beans.Logger.Fatal(err)
				ch <- os.Interrupt
			}
		}()
	}))
	defer func() {
		panicOnErr(c.Invoke(func(srv *api.Server) {
			if err := srv.Shutdown(); err != nil {
				beans.Logger.Fatal(err)
			}
		}))
	}()

	<-ch
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
