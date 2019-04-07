package main

import (
	"fmt"
	"os"
	"os/signal"

	"go.uber.org/dig"

	"github.com/studtool/users-service/api"
	"github.com/studtool/users-service/beans"
	"github.com/studtool/users-service/config"
)

func main() {
	c := dig.New()

	panicOnErr(c.Provide(api.NewServer))

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)

	_ = c.Invoke(func(srv *api.Server) {
		go func() {
			if err := srv.Run(); err != nil {
				beans.Logger.Fatal(err)
				ch <- os.Interrupt
			}
		}()

		beans.Logger.Info(fmt.Sprintf("server: started; [port: %s]", config.ServerPort.Value()))
	})
	defer func() {
		_ = c.Invoke(func(srv *api.Server) {
			if err := srv.Shutdown(); err != nil {
				beans.Logger.Fatal(err)
			} else {
				beans.Logger.Info("server: down")
			}
		})
	}()

	<-ch
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
