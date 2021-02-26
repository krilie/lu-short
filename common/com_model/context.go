package com_model

import "context"

// AppCtxValuesName app 层面上下文信息
var AppCtxValuesName = "AppCtxValuesName"

type AppCtxValues struct {
	TrackId  string
	AppName  string
	HostName string
}

func MustGetCtxValues(ctx context.Context) *AppCtxValues {
	value := ctx.Value(AppCtxValuesName)
	if value == nil {
		panic("get app ctx values error,no ctx value.")
	}
	return value.(*AppCtxValues)
}

func NewCtxWithValues(ctx context.Context, val *AppCtxValues) context.Context {
	return context.WithValue(ctx, AppCtxValuesName, val)
}
