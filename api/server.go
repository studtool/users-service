package api

import (
	"fmt"
	"net/http"

	"go.uber.org/dig"

	"github.com/go-http-utils/headers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/studtool/common/logs"
	"github.com/studtool/common/rest"

	"github.com/studtool/users-service/config"
	"github.com/studtool/users-service/repositories"
	"github.com/studtool/users-service/utils"
)

type Server struct {
	rest.Server

	structLogger  logs.Logger
	reflectLogger logs.Logger

	usersRepository repositories.UsersRepository
}

type ServerParams struct {
	dig.In

	UsersRepository repositories.UsersRepository
}

func NewServer(params ServerParams) *Server {
	srv := &Server{
		usersRepository: params.UsersRepository,
	}

	v := rest.ParseAPIVersion(config.ComponentVersion)
	srvPublicPath := rest.MakeAPIPath(v, rest.APITypePublic, "/users")
	srvProtectedPath := rest.MakeAPIPath(v, rest.APITypeProtected, "/users")

	mx := mux.NewRouter()
	mx.Handle(srvPublicPath, handlers.MethodHandler{
		http.MethodGet: http.HandlerFunc(srv.findProfile),
	})
	mx.Handle(srvPublicPath+"/{user_id}/profile", handlers.MethodHandler{
		http.MethodGet: http.HandlerFunc(srv.getProfile),
	})
	mx.Handle(srvProtectedPath+"/{user_id}/profile", handlers.MethodHandler{
		http.MethodPatch: srv.WithAuth(http.HandlerFunc(srv.updateProfile)),
	})
	mx.Handle(`/pprof`, rest.GetProfilerHandler())
	mx.Handle(`/metrics`, rest.GetMetricsHandler())

	reqHandler := srv.WithRecover(mx)
	if config.RequestsLogsEnabled.Value() {
		reqHandler = srv.WithLogs(reqHandler)
	}
	if config.CorsAllowed.Value() {
		reqHandler = srv.WithCORS(reqHandler, rest.CORS{
			Origins: []string{"*"},
			Methods: []string{
				http.MethodGet, http.MethodHead,
				http.MethodPost, http.MethodPatch,
				http.MethodDelete, http.MethodOptions,
			},
			Headers: []string{
				headers.Authorization, headers.UserAgent,
				headers.ContentType, headers.ContentLength,
				headers.ContentEncoding, headers.ContentLanguage,
			},
			Credentials: false,
		})
	}

	srv.structLogger = srvutils.MakeStructLogger(srv)
	srv.reflectLogger = srvutils.MakeReflectLogger(srv)

	srv.Server = *rest.NewServer(
		rest.ServerParams{
			Address: fmt.Sprintf(":%d", config.ServerPort.Value()),
			Handler: reqHandler,

			StructLogger:  srv.structLogger,
			ReflectLogger: srv.reflectLogger,
			RequestLogger: srvutils.MakeRequestLogger(srv),

			APIClassifier: rest.NewPathAPIClassifier(),
		},
	)
	srv.structLogger.Info("initialized")

	return srv
}
