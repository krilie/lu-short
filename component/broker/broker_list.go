package broker

import (
	go_smq "github.com/krilie/go-smq"
	"lu-short/component/broker/messages"
	"reflect"
)

type Brokers struct {
	MsgRedirect *go_smq.Broker
}

func NewBrokers() *Brokers {
	return &Brokers{
		MsgRedirect: go_smq.NewStartedBroker(reflect.TypeOf(&messages.MsgRedirect{}), 200),
	}
}
