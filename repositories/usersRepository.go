package repositories

import (
	"github.com/studtool/common/errs"

	"github.com/studtool/users-service/models"
)

type UsersRepository interface {
	AddUserById(userId string) *errs.Error
	FindUserInfoByUsername(u *models.UserInfo) *errs.Error
	GetUser(u *models.User) *errs.Error
	UpdateUser(u *models.User) *errs.Error
	DeleteUserById(userId string) *errs.Error
}
