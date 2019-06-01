package main

import (
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/dig"

	"github.com/studtool/common/logs"
	"github.com/studtool/common/utils/assertions"

	"github.com/studtool/users-service/api"
	"github.com/studtool/users-service/config"
	"github.com/studtool/users-service/messages"
	"github.com/studtool/users-service/repositories"
	"github.com/studtool/users-service/repositories/mongo"
)

func main() {
	c := dig.New()
	logger := logs.NewRawLogger()

	if config.RepositoriesEnabled.Value() {
		assertions.AssertOk(c.Provide(mongo.NewConnection))
		assertions.AssertOk(c.Invoke(func(conn *mongo.Connection) {
			if err := conn.Open(); err != nil {
				logger.Fatal(err.Error())
			} else {
				logger.Info("storage: connection open")
			}
		}))
		defer func() {
			assertions.AssertOk(c.Invoke(func(conn *mongo.Connection) {
				if err := conn.Close(); err != nil {
					logger.Fatal(err)
				} else {
					logger.Info("storage: connection closed")
				}
			}))
		}()

		assertions.AssertOk(c.Provide(
			mongo.NewUsersRepository,
			dig.As(new(repositories.UsersRepository)),
		))
	} else {
		assertions.AssertOk(c.Provide(
			func() repositories.UsersRepository {
				return nil
			},
		))
	}

	if config.QueuesEnabled.Value() {
		assertions.AssertOk(c.Provide(messages.NewMqClient))
		assertions.AssertOk(c.Invoke(func(q *messages.MqClient) {
			if err := q.OpenConnection(); err != nil {
				logger.Fatal(err)
			} else {
				logger.Info("queue: connection open")
			}
			if err := q.Run(); err != nil {
				logger.Fatal(err)
			} else {
				logger.Info("queue: ready to receive messages")
			}
		}))
		defer func() {
			assertions.AssertOk(c.Invoke(func(q *messages.MqClient) {
				if err := q.CloseConnection(); err != nil {
					logger.Fatal(err)
				} else {
					logger.Info("queue: connection closed")
				}
			}))
		}()
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT)
	signal.Notify(ch, syscall.SIGTERM)

	assertions.AssertOk(c.Provide(api.NewServer))
	assertions.AssertOk(c.Invoke(func(srv *api.Server) {
		go func() {
			if err := srv.Run(); err != nil {
				logger.Fatal(err)
				ch <- os.Interrupt
			}
		}()
	}))
	defer func() {
		assertions.AssertOk(c.Invoke(func(srv *api.Server) {
			if err := srv.Shutdown(); err != nil {
				logger.Fatal(err)
			}
		}))
	}()

	<-ch
}
