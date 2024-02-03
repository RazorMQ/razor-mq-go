package pc

type ClientConsumer struct {
	SubscribedTopics []string
	Channel          *chan []byte
}
