package messages

import (
	"context"
	"reflect"
)

type TestMessage struct {
	Test string
	ctx  context.Context
}

func (t *TestMessage) GetName() string {
	return "for test"
}

func (t *TestMessage) GetCtx() context.Context {
	return t.ctx
}

func (t *TestMessage) GetType() reflect.Type {
	return reflect.TypeOf(t)
}
