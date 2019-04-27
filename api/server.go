package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/studtool/common/consts"
	"github.com/studtool/common/rest"

	"github.com/studtool/users-service/beans"
	"github.com/studtool/users-service/config"
	"github.com/studtool/users-service/repositories"
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
	mx.Handle(fmt.Sprintf("/api/users/{%s}/profile", userIdVar), handlers.MethodHandler{
		http.MethodGet:   http.HandlerFunc(srv.getProfile),
		http.MethodPatch: srv.server.WithAuth(http.HandlerFunc(srv.updateProfile)),
	})

	srv.server.SetLogger(beans.Logger)

	h := srv.server.WithRecover(mx)
	if config.ShouldLogRequests.Value() {
		h = srv.server.WithLogs(h)
	}
	srv.server.SetHandler(h)

	return srv
}

func (srv *Server) Run() error {
	return srv.server.Run()
}

func (srv *Server) Shutdown() error {
	return srv.server.Shutdown()
}
