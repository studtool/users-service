package api

import (
	"github.com/studtool/common/consts"
	"github.com/studtool/common/errs"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	userIdVar     = `user_id`
	userIdPattern = `\w{8}-\w{4}-\w{4}-\w{4}-\w{12}`
)

func (srv *Server) parseUsername(r *http.Request) (string, *errs.Error) {
	v := r.URL.Query().Get("username")
	if v == consts.EmptyString {
		return consts.EmptyString, errs.NewBadFormatError(
			`no "username" query parameter`,
		)
	}
	return v, nil
}

func (srv *Server) parseUserId(r *http.Request) string {
	return mux.Vars(r)[userIdVar]
}

func (srv *Server) checkAuthPermission(r *http.Request) *errs.Error {
	if srv.server.ParseUserId(r) == srv.parseUserId(r) {
		return nil
	}
	return errs.NewPermissionDeniedError("access denied") //TODO save err
}
