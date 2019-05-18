package messages

import (
	"fmt"

	"github.com/mailru/easyjson"
	"github.com/streadway/amqp"

	"github.com/studtool/common/consts"
	"github.com/studtool/common/errs"

	"github.com/studtool/users-service/beans"
)

func (c *MqClient) declareQueue(queueName string) (amqp.Queue, error) {
	return c.channel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
}

func (c *MqClient) runConsumer(
	queueName string,
	handler messageHandler,
) error {
	messages, err := c.channel.Consume(
		queueName,
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
			handler(d.Body)
		}
	}()

	return nil
}

func (c *MqClient) marshalMessageBody(v easyjson.Marshaler) ([]byte, error) {
	return easyjson.Marshal(v)
}

func (c *MqClient) unmarshalMessageBody(data []byte, v easyjson.Unmarshaler) error {
	return easyjson.Unmarshal(data, v)
}

func (c *MqClient) writeMessageConsumedLog(
	queueName string,
	data easyjson.Marshaler,
) {
	beans.Logger().Info(
		fmt.Sprintf("message consumed (%s): %v", queueName, data),
	)
}

func (c *MqClient) writeMessageConsumptionErrorLog(
	queueName string,
	data easyjson.Marshaler,
) {
	beans.Logger().Error(
		fmt.Sprintf("message not consumed (%s): %v", queueName, data),
	)
}

func (c *MqClient) writeErrorLog(err error) {
	if err != nil {
		beans.Logger().Error(err)
	}
}

func (c *MqClient) writeRepositoryErrorLog(err *errs.Error) {
	if err != nil {
		beans.Logger().Error(err)
	}
}
