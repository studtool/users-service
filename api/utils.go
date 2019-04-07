package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	userIdVar     = `user_id`
	userIdPattern = `\w{8}-\w{4}-\w{4}-\w{4}-\w{12}`
)

func (srv *Server) parseUserId(r *http.Request) string {
	return mux.Vars(r)[userIdVar]
}
