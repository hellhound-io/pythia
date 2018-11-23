package nats

import (
	"encoding/json"
	"fmt"
	xnats "github.com/nats-io/go-nats"
	syslog "github.com/prometheus/common/log"
	"github.com/sirupsen/logrus"
	"gitlab.com/consensys-hellhound/pythia/domain/node"
	"gitlab.com/consensys-hellhound/pythia/log"
)

type broker struct {
	url string
	nc  *xnats.Conn
}

func NewBroker(url string) node.Broker {
	b := &broker{
		url: url,
	}
	go b.init()
	return b
}

func (b *broker) init() {
	var err error
	b.nc, err = xnats.Connect(b.url)
	if err != nil {
		syslog.Fatalf("cannot connect to NATS : %s", err.Error())
	}
}

func (b broker) Publish(topic string, event interface{}) (err error) {
	payload, err := json.Marshal(event)
	if err != nil {
		return
	}
	log.Logger.WithFields(logrus.Fields{
		log.Method: "nats.publish",
		"topic":    topic,
	}).Infoln(string(payload))
	err = b.nc.Publish(topic, payload)
	return
}

func (b broker) Subscribe(topic string, handler func([]byte) error) error {
	_, err := b.nc.Subscribe(topic, func(m *xnats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
		handler(m.Data)
	})
	return err
}
