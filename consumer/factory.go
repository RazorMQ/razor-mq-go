package consumer

import (
	"github.com/RazorMQ/razor-mq-go/client"
	"github.com/RazorMQ/razor-mq-go/message"
)

type ConsumerSettings struct {
	SubscribedTopics []string
	IgnoreOldMessage bool
	Handler          func(message.EnqueuedMessage)
}

type Factory interface {
	NewConsumer(*client.RazorMQClient, ConsumerSettings) *Consumer
}
