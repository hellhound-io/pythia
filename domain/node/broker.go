package node

type Broker interface {
	Publish(topic string, event interface{}) error
	Subscribe(topic string, handler func([]byte) error) error
}
