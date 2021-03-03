package broker

import (
	"context"
	"github.com/stretchr/testify/assert"
	"lu-short/component/broker/messages"
	"testing"
)

func TestAutoFunc(t *testing.T) {
	Smq := NewBroker()
	count := 0
	// register smq
	err := Smq.Register(context.Background(), func(i *messages.TestMessage) {
		println(i.Test)
		count++
		assert.Equal(t, "for test test", i.Test, "should equal")
	})
	assert.Equal(t, nil, err, "should nil")
	// send message
	for i := 0; i < 777; i++ {
		err = Smq.Send(context.Background(), &messages.TestMessage{Test: "for test test"})
		assert.Equal(t, nil, err, "should nil")
		if err != nil {
			t.FailNow()
		}
	}
	// check send message
	Smq.Close()
	assert.Equal(t, 777, count, "no match count")
}
