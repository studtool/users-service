package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/studtool/users-service/config"
)

const (
	ConnectTimeout    = 5 * time.Second
	DisconnectTimeout = 5 * time.Second
)

type Connection struct {
	uri    string
	client *mongo.Client
}

func NewConnection() *Connection {
	return &Connection{
		uri: fmt.Sprintf("mongodb://%s:%s",
			config.StorageHost.Value(), config.StoragePort.Value(),
		),
	}
}

func (conn *Connection) Open() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(conn.uri))
	if err != nil {
		return err
	}

	err = client.Connect(conn.tdContext(ConnectTimeout))
	if err != nil {
		return err
	}

	conn.client = client
	return nil
}

func (conn *Connection) Close() error {
	return conn.client.Disconnect(conn.tdContext(DisconnectTimeout))
}

func (conn *Connection) tdContext(t time.Duration) context.Context {
	ctx, _ := context.WithTimeout(context.TODO(), t)
	return ctx
}

func (conn *Connection) bgContext(t time.Duration) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), t)
	return ctx
}
