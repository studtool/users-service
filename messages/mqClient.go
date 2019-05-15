package messages

import (
	"fmt"

	"github.com/mailru/easyjson"
	"github.com/streadway/amqp"
	"go.uber.org/dig"

	"github.com/studtool/common/consts"
	"github.com/studtool/common/errs"
	"github.com/studtool/common/queues"
	"github.com/studtool/common/utils"

	"github.com/studtool/users-service/beans"
	"github.com/studtool/users-service/config"
	"github.com/studtool/users-service/repositories"
)

type QueueClient struct {
	connStr    string
	connection *amqp.Connection

	channel *amqp.Channel

	createdUsersQueue amqp.Queue
	deletedUsersQueue amqp.Queue

	usersRepository repositories.UsersRepository
}

type ClientParams struct {
	dig.In

	repositories.UsersRepository
}

func NewQueueClient(params ClientParams) *QueueClient {
	return &QueueClient{
		connStr: fmt.Sprintf("amqp://%s:%s@%s:%d/",
			config.UsersMqUser.Value(), config.UsersMqPassword.Value(),
			config.UsersMqHost.Value(), config.UsersMqPort.Value(),
		),
		usersRepository: params.UsersRepository,
	}
}

func (c *QueueClient) OpenConnection() error {
	var conn *amqp.Connection
	err := utils.WithRetry(func(n int) (err error) {
		if n > 0 {
			beans.Logger.Info(fmt.Sprintf("opening message queue connection. retry #%d", n))
		}
		conn, err = amqp.Dial(c.connStr)
		return err
	}, config.UsersMqConnNumRet.Value(), config.UsersMqConnRetItv.Value())
	if err != nil {
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	c.createdUsersQueue, err = ch.QueueDeclare(
		queues.CreatedUsersQueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	c.deletedUsersQueue, err = ch.QueueDeclare(
		queues.DeletedUsersQueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	c.channel = ch
	c.connection = conn

	return nil
}

func (c *QueueClient) CloseConnection() error {
	if err := c.channel.Close(); err != nil {
		return err
	}
	return c.connection.Close()
}

type MessageHandler func(data []byte)

func (c *QueueClient) Run() error {
	if err := c.recvCreatedUsersData(); err != nil {
		return err
	}
	if err := c.recvDeletedUsersData(); err != nil {
		return err
	}
	return nil
}

func (c *QueueClient) recvCreatedUsersData() error {
	messages, err := c.channel.Consume(
		c.createdUsersQueue.Name,
		consts.EmptyString,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range messages {
			c.addUser(d.Body)
		}
	}()

	return nil
}

func (c *QueueClient) recvDeletedUsersData() error {
	messages, err := c.channel.Consume(
		c.deletedUsersQueue.Name,
		consts.EmptyString,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range messages {
			c.deleteUser(d.Body)
		}
	}()

	return nil
}

func (c *QueueClient) parseMessageBody(data []byte, v easyjson.Unmarshaler) error {
	return easyjson.Unmarshal(data, v)
}

func (c *QueueClient) addUser(body []byte) {
	data := &queues.CreatedUserData{}
	if err := c.parseMessageBody(body, data); err != nil {
		c.handleErr(err)
	} else {
		if err := c.usersRepository.AddUserById(data.UserID); err != nil {
			c.handleRepoErr(err)
		}
	}
}

func (c *QueueClient) deleteUser(body []byte) {
	data := &queues.DeletedUserData{}
	if err := c.parseMessageBody(body, data); err != nil {
		c.handleErr(err)
	} else {
		if err := c.usersRepository.DeleteUserById(data.UserID); err != nil {
			c.handleRepoErr(err)
		}
	}
}

func (c *QueueClient) handleErr(err error) {
	if err != nil {
		beans.Logger.Error(err)
	}
}

func (c *QueueClient) handleRepoErr(err *errs.Error) {
	if err != nil {
		beans.Logger.Error(err)
	}
}
