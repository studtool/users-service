package mq

import (
	"fmt"

	"github.com/streadway/amqp"

	"github.com/studtool/common/consts"
	"github.com/studtool/common/errs"
	"github.com/studtool/common/utils"

	"github.com/studtool/users-service/beans"
	"github.com/studtool/users-service/config"
	"github.com/studtool/users-service/repositories"
)

type MQ struct {
	connStr string

	ch   *amqp.Channel
	conn *amqp.Connection

	cq amqp.Queue
	dq amqp.Queue

	usersRepository repositories.UsersRepository
}

func NewQueue(uRepo repositories.UsersRepository) *MQ {
	return &MQ{
		connStr: fmt.Sprintf("amqp://%s:%s@%s:%s/",
			config.UsersMqUser.Value(), config.UsersMqPassword.Value(),
			config.UsersMqHost.Value(), config.UsersMqPort.Value(),
		),
		usersRepository: uRepo,
	}
}

func (mq *MQ) OpenConnection() error {
	var conn *amqp.Connection
	err := utils.WithRetry(func(n int) (err error) {
		if n > 0 {
			beans.Logger.Info(fmt.Sprintf("opening message queue connection. retry #%d", n))
		}
		conn, err = amqp.Dial(mq.connStr)
		return err
	}, config.UsersMqConnNumRet.Value(), config.UsersMqConnRetItv.Value())
	if err != nil {
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	mq.cq, err = ch.QueueDeclare(
		config.CreatedUsersQueueName.Value(),
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	mq.dq, err = ch.QueueDeclare(
		config.DeletedUsersQueueName.Value(),
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	mq.ch = ch
	mq.conn = conn

	return nil
}

func (mq *MQ) CloseConnection() error {
	if err := mq.ch.Close(); err != nil {
		return err
	}
	return mq.conn.Close()
}

func (mq *MQ) Run() error {
	if err := mq.receive(mq.cq, mq.addUser); err != nil {
		return err
	}
	return nil
}

type MessageHandler func(userId string) *errs.Error

func (mq *MQ) receive(q amqp.Queue, h MessageHandler) error {
	messages, err := mq.ch.Consume(
		q.Name,
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
			if err := h(string(d.Body)); err != nil {
				beans.Logger.Error(err)
			}
		}
	}()

	return nil
}

func (mq *MQ) addUser(userId string) *errs.Error {
	return mq.usersRepository.AddUserById(userId) //TODO should retry if failed
}
