package messages

import (
	"fmt"

	"github.com/mailru/easyjson"
	"github.com/streadway/amqp"

	"github.com/studtool/common/consts"
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
	queueName string, data easyjson.Marshaler,
) {
	c.structLogger.Infof(
		"message consumed (%s): %v", queueName, data,
	)
}

func (c *MqClient) writeMessageConsumptionErrorLog(
	queueName string, data easyjson.Marshaler,
) {
	c.structLogger.Errorf(
		fmt.Sprintf("message not consumed (%s): %v", queueName, data),
	)
}
