package api

import (
	"net/http"
)

func (srv *Server) findProfile(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func (srv *Server) getProfile(w http.ResponseWriter, r *http.Request) {
	userId := srv.parseUserId(r)

	m, err := srv.usersRepository.GetUser(userId)
	if err != nil {
		srv.server.WriteErrJSON(w, err)
		return
	}

	srv.server.WriteBodyJSON(w, http.StatusOK, m)
}

func (srv *Server) updateProfile(w http.ResponseWriter, r *http.Request) {
	//TODO
}
