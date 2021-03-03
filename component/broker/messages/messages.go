package messages

import (
	"context"
	"reflect"
	"time"
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

// 跳转发生
type MsgRedirect struct {
	Ctx         context.Context
	RedirectId  string
	CustomerId  string
	OriUrl      string
	ShortUrl    string
	TrackId     string
	Device      string
	RemoteIp    string
	VisitorTime time.Time
}

func (m *MsgRedirect) GetName() string {
	return "MsgRedirect"
}

func (m *MsgRedirect) GetCtx() context.Context {
	return m.Ctx
}

func (m *MsgRedirect) GetType() reflect.Type {
	return reflect.TypeOf(m)
}
