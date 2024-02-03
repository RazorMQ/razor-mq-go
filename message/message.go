package message

type EnqueuedMessage struct {
	ProducerHost string
	Topic        string
	Timestamp    string
	Index        int64
	Data         []byte
}

type Message struct {
	Data  []byte
	Topic string
}
