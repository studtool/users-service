package mongo

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/studtool/common/errs"

	"github.com/studtool/users-service/config"
	"github.com/studtool/users-service/models"
)

const (
	InsertTimeout = time.Second
	SelectTimeout = time.Second
	DeleteTimeout = time.Second
)

type UsersRepository struct {
	connection *Connection
	collection *mongo.Collection
}

func NewUsersRepository(conn *Connection) *UsersRepository {
	db := conn.client.Database(config.StorageDB.Value())
	return &UsersRepository{
		collection: db.Collection("users"),
	}
}

func (r *UsersRepository) AddUserById(userId string) *errs.Error {
	ctx := r.connection.bgContext(InsertTimeout)
	_, err := r.collection.InsertOne(ctx, bson.M{
		"_id":      userId,
		"username": fmt.Sprintf("user%s", userId),
	})
	if err != nil {
		return r.wrapErr(err)
	}
	return nil
}

func (r *UsersRepository) FindUserInfoByUsername(u *models.UserInfo) *errs.Error {
	ctx := r.connection.tdContext(SelectTimeout)
	res := r.collection.FindOne(ctx, bson.M{
		"username": u.Username,
	})
	if err := res.Err(); err != nil {
		panic(err) //TODO parse err
	}

	var m bson.M
	if err := res.Decode(&m); err != nil {
		return r.wrapErr(err) //TODO check empty?
	}

	userId, ok := m["userId"]
	if !ok {
		return errs.NewInternalError(
			fmt.Sprintf(`storage: %v should contain "userId"`, m),
		)
	}

	if v, ok := userId.(string); ok {
		u.Id = v
	} else {
		return errs.NewInternalError(
			`storage: "userId" should be string`, //TODO err to struct
		)
	}

	return nil
}

func (r *UsersRepository) GetUser(userId string) (*models.UserMap, *errs.Error) {
	ctx := r.connection.tdContext(SelectTimeout)
	res := r.collection.FindOne(ctx, bson.M{
		"_id": userId,
	})
	if err := res.Err(); err != nil {
		return nil, r.wrapErr(err)
	}

	m := &bson.M{}
	if err := res.Decode(&m); err != nil {
		return nil, r.parseErr(err)
	}

	(*m)["userId"] = (*m)["_id"]
	delete(*m, "_id")

	return (*models.UserMap)(m), nil
}

func (r *UsersRepository) UpdateUser(u *models.User) *errs.Error {
	panic("implement me") //TODO
}

func (r *UsersRepository) DeleteUserById(userId string) *errs.Error {
	ctx := r.connection.tdContext(DeleteTimeout)
	_, err := r.collection.DeleteOne(ctx, bson.M{
		"_id": userId,
	})
	if err != nil {
		return r.wrapErr(err)
	}
	return nil
}

func (r *UsersRepository) parseErr(err error) *errs.Error {
	if err.Error() == "mongo: no documents in result" {
		return errs.NewNotFoundError("user not found")
	}
	return r.wrapErr(err)
}

func (r *UsersRepository) wrapErr(err error) *errs.Error {
	return errs.NewInternalError(err.Error())
}

func (r *UsersRepository) GetComponent() string {
	return Component
}
