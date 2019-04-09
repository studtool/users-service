package mongo

import (
	"context"
	"fmt"
	"github.com/studtool/structs"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/studtool/common/errs"

	"github.com/studtool/users-service/config"
	"github.com/studtool/users-service/models"
)

const (
	SelectTimeout = time.Second
	UpdateTimeout = 2 * time.Second
)

const (
	idField = "_id"
)

type UsersRepository struct {
	connection  *Connection
	collection  *mongo.Collection
	notFoundErr *errs.Error
}

func NewUsersRepository(conn *Connection) *UsersRepository {
	db := conn.client.Database(config.StorageDB.Value())
	return &UsersRepository{
		connection:  conn,
		collection:  db.Collection("users"),
		notFoundErr: errs.NewNotFoundError("profile not found"),
	}
}

func (r *UsersRepository) AddUserById(userId string) *errs.Error {
	_, err := r.collection.InsertOne(context.TODO(), bson.M{
		idField:    userId,
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
		return r.parseErr(err)
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
		idField: userId,
	})
	if err := res.Err(); err != nil {
		return nil, r.wrapErr(err)
	}

	m := &bson.M{}
	if err := res.Decode(&m); err != nil {
		return nil, r.parseErr(err)
	}

	(*m)["userId"] = (*m)[idField]
	delete(*m, idField)

	return (*models.UserMap)(m), nil
}

func (r *UsersRepository) UpdateUser(u *models.User) *errs.Error {
	m := structs.Map(u)

	ctx := r.connection.tdContext(UpdateTimeout)
	res, err := r.collection.UpdateOne(ctx, bson.M{
		"_id": u.Id,
	}, bson.D{{"$set", bson.M(m)}})
	if err != nil {
		return r.wrapErr(err)
	}
	if res.MatchedCount != 1 {
		return r.notFoundErr
	}

	return nil
}

func (r *UsersRepository) DeleteUserById(userId string) *errs.Error {
	_, err := r.collection.DeleteOne(context.TODO(), bson.M{
		"_id": userId,
	})
	if err != nil {
		return r.wrapErr(err)
	}
	return nil
}

func (r *UsersRepository) parseErr(err error) *errs.Error {
	if err.Error() == "mongo: no documents in result" {
		return r.notFoundErr
	}
	return r.wrapErr(err)
}

func (r *UsersRepository) wrapErr(err error) *errs.Error {
	return errs.NewInternalError(err.Error())
}
