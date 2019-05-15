package api

import (
	"github.com/studtool/users-service/models"
	"net/http"
)

func (srv *Server) findProfile(w http.ResponseWriter, r *http.Request) {
	username, err := srv.parseUsername(r)
	if err != nil {
		srv.server.WriteErrJSON(w, err)
		return
	}

	user := &models.UserInfo{
		Username: username,
	}
	if err := srv.usersRepository.FindUserInfoByUsername(user); err != nil {
		srv.server.WriteErrJSON(w, err)
		return
	}

	srv.server.WriteOkJSON(w, user)
}

func (srv *Server) getProfile(w http.ResponseWriter, r *http.Request) {
	userId := srv.parseUserId(r)

	m, err := srv.usersRepository.GetUser(userId)
	if err != nil {
		srv.server.WriteErrJSON(w, err)
		return
	}

	srv.server.WriteOkJSON(w, m)
}

func (srv *Server) updateProfile(w http.ResponseWriter, r *http.Request) {
	userId := srv.parseUserId(r)
	if err := srv.checkAuthPermission(r); err != nil {
		srv.server.WriteErrJSON(w, err)
		return
	}

	user := &models.User{}
	if err := srv.server.ParseBodyJSON(user, r); err != nil {
		srv.server.WriteErrJSON(w, err)
		return
	}

	user.Id = userId
	if err := srv.usersRepository.UpdateUser(user); err != nil {
		srv.server.WriteErrJSON(w, err)
		return
	}

	srv.server.WriteOk(w)
}
