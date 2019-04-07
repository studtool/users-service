package api

import (
	"fmt"
	"github.com/studtool/users-service/repositories"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/studtool/common/consts"
	"github.com/studtool/common/rest"

	"github.com/studtool/users-service/beans"
	"github.com/studtool/users-service/config"
)

type Server struct {
	server          *rest.Server
	usersRepository repositories.UsersRepository
}

func NewServer(r repositories.UsersRepository) *Server {
	srv := &Server{
		server: rest.NewServer(
			rest.ServerConfig{
				Host: consts.EmptyString,
				Port: config.ServerPort.Value(),
			},
		),
		usersRepository: r,
	}

	mx := mux.NewRouter()
	mx.Handle("/api/users", handlers.MethodHandler{
		http.MethodGet: http.HandlerFunc(srv.findProfile),
	})
	mx.Handle(fmt.Sprintf("/api/users/{%s:%s}/profile", userIdVar, userIdPattern), handlers.MethodHandler{
		http.MethodGet:   http.HandlerFunc(srv.getProfile),
		http.MethodPatch: srv.server.WithAuth(http.HandlerFunc(srv.updateProfile)),
	})

	srv.server.SetLogger(beans.Logger)
	srv.server.SetHandler(srv.server.WithRecover(mx))

	return srv
}

func (srv *Server) Run() error {
	return srv.server.Run()
}

func (srv *Server) Shutdown() error {
	return srv.server.Shutdown()
}
