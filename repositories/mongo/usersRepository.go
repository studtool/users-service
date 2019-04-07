package mongo

import (
	"github.com/studtool/common/errs"

	"github.com/studtool/users-service/models"
)

type UsersRepository struct {
}

func NewUsersRepository(conn *Connection) *UsersRepository {
	return &UsersRepository{}
}

func (*UsersRepository) AddUserById(userId string) *errs.Error {
	panic("implement me") //TODO
}

func (*UsersRepository) FindUserInfoByUsername(u *models.UserInfo) *errs.Error {
	panic("implement me") //TODO
}

func (*UsersRepository) GetUser(u *models.User) *errs.Error {
	panic("implement me") //TODO
}

func (*UsersRepository) UpdateUser(u *models.User) *errs.Error {
	panic("implement me") //TODO
}

func (*UsersRepository) DeleteUserById(userId string) *errs.Error {
	panic("implement me") //TODO
}
