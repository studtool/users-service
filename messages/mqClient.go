package messages

import (
	"fmt"

	"go.uber.org/dig"

	"github.com/streadway/amqp"

	"github.com/studtool/common/logs"
	"github.com/studtool/common/queues"
	"github.com/studtool/common/utils"

	"github.com/studtool/users-service/config"
	"github.com/studtool/users-service/repositories"
	"github.com/studtool/users-service/utils"
)

type MqClient struct {
	connStr    string
	connection *amqp.Connection

	channel *amqp.Channel

	profilesToCreateQueue amqp.Queue
	profilesToDeleteQueue amqp.Queue

	usersRepository repositories.UsersRepository

	structLogger  logs.Logger
	reflectLogger logs.Logger
}

type MqClientParams struct {
	dig.In

	repositories.UsersRepository
}

func NewMqClient(params MqClientParams) *MqClient {
	c := &MqClient{
		connStr: fmt.Sprintf("amqp://%s:%s@%s:%d/",
			config.MqUser.Value(), config.MqPassword.Value(),
			config.MqHost.Value(), config.MqPort.Value(),
		),
		usersRepository: params.UsersRepository,
	}

	c.structLogger = srvutils.MakeStructLogger(c)
	c.reflectLogger = srvutils.MakeReflectLogger(c)

	return c
}

func (c *MqClient) OpenConnection() error {
	var conn *amqp.Connection
	err := utils.WithRetry(func(n int) (err error) {
		if n > 0 {
			c.structLogger.Infof("opening message queue connection. retry #%d", n)
		}
		conn, err = amqp.Dial(c.connStr)
		return err
	}, config.MqConnNumRet.Value(), config.MqConnRetItv.Value())
	if err != nil {
		return err
	}

	c.connection = conn

	c.channel, err = conn.Channel()
	if err != nil {
		return err
	}

	c.profilesToCreateQueue, err =
		c.declareQueue(queues.ProfilesToCreateQueueName)
	if err != nil {
		return err
	}

	c.profilesToDeleteQueue, err =
		c.declareQueue(queues.ProfilesToDeleteQueueName)
	if err != nil {
		return err
	}

	return nil
}

func (c *MqClient) CloseConnection() error {
	if err := c.channel.Close(); err != nil {
		return err
	}
	return c.connection.Close()
}

type MessageHandler func(data []byte)

func (c *MqClient) Run() error {
	err := c.runConsumer(
		queues.ProfilesToCreateQueueName,
		c.createProfile,
	)
	if err != nil {
		return err
	}

	err = c.runConsumer(
		queues.ProfilesToDeleteQueueName,
		c.deleteProfile,
	)
	if err != nil {
		return err
	}

	return nil
}
