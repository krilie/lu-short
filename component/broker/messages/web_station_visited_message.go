package messages

import (
	"context"
	"reflect"
	"time"
)

// WebStationVisitedMessage 网站被打开的消息
type WebStationVisitedMessage struct {
	Ctx        context.Context
	AccessTime time.Time
	Ip         string
	TraceId    string
}

func (w *WebStationVisitedMessage) GetName() string {
	return "WebStationVisitedMessage"
}

func (w *WebStationVisitedMessage) GetCtx() context.Context {
	return w.Ctx
}

func (w *WebStationVisitedMessage) GetType() reflect.Type {
	return reflect.TypeOf(w)
}
