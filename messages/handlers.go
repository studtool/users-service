package messages

import (
	"github.com/studtool/common/queues"
)

type messageHandler func(data []byte)

func (c *MqClient) createProfile(body []byte) {
	data := &queues.ProfileToCreateData{}
	if err := c.unmarshalMessageBody(body, data); err != nil {
		c.writeErrorLog(err)
	} else {
		if err := c.usersRepository.AddUserById(data.UserID); err != nil {
			c.writeRepositoryErrorLog(err)
		} else {
			c.writeMessageConsumedLog(
				queues.ProfilesToCreateQueueName, data,
			)
		}
	}
}

func (c *MqClient) deleteProfile(body []byte) {
	data := &queues.ProfileToDeleteData{}
	if err := c.unmarshalMessageBody(body, data); err != nil {
		c.writeErrorLog(err)
	} else {
		if err := c.usersRepository.DeleteUserById(data.UserID); err != nil {
			c.writeRepositoryErrorLog(err)
		} else {
			c.writeMessageConsumedLog(
				queues.ProfilesToDeleteQueueName, data,
			)
		}
	}
}
