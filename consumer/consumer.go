package consumer

import "github.com/RazorMQ/razor-mq-go/message"

type ConsumerInterface interface {
	HandleMessage([]byte)
	HandleDisconnect()
	GetLastReadIndex() int
}

type Consumer struct {
	lastReadIndex     int
	IgnoreOldMessages bool
	handler           func(message.EnqueuedMessage)
}

func (c *Consumer) HandleMessage(msg message.EnqueuedMessage) {
	if c.IgnoreOldMessages && c.lastReadIndex < c.lastReadIndex {
		return
	}
	c.lastReadIndex = int(msg.Index)
	c.handler(msg)
}

func (c *Consumer) HandleDisconnect() {

}

func (c *Consumer) GetLastReadIndex() int {
	return c.lastReadIndex
}

func (c *Consumer) StartListening() {

}
