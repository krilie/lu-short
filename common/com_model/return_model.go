package com_model

import (
	"lu-short/common/errs"
)

var StdSuccess = &CommonReturn{
	Code:    errs.Success.ToInt(),
	Message: "success",
	Detail:  nil,
	Data:    nil,
}

// CommonReturn
type CommonReturn struct {
	Code    int         `json:"code" swaggo:"true,错误码" example:"400"`
	Message string      `json:"message" swaggo:"true,错误信息" example:"错误信息"`
	Detail  *string     `json:"detail,omitempty" swaggo:"false,错误的详细信息，用于排查错误"  example:"错误的详细信息，用于排查错误"` // 可由运行模式控制是否显示
	Data    interface{} `json:"data,omitempty" `                                                          // 数据值
}

func NewRet(err *errs.Err) *CommonReturn {
	fullMsg := err.GetFullMsg()
	return &CommonReturn{
		Code:    err.Code.ToInt(),
		Message: err.Message,
		Detail:  &fullMsg,
		Data:    nil,
	}
}
func NewRetFromErr(err error) *CommonReturn {
	fullMsg := errs.GetErrMsg(err)
	nErr, ok := err.(*errs.Err)
	if ok {
		return NewRet(nErr)
	} else {
		return &CommonReturn{
			Code:    errs.ErrorNormal.ToInt(),
			Message: err.Error(),
			Detail:  &fullMsg,
			Data:    nil,
		}
	}
}
func NewFailure(code errs.ErrCode, msg string) *CommonReturn {
	if msg == "" {
		msg = "failure"
	}
	return &CommonReturn{
		Code:    code.ToInt(),
		Message: msg,
		Detail:  nil,
		Data:    nil,
	}
}

func NewSuccess(data interface{}) *CommonReturn {
	return &CommonReturn{
		Code:    errs.Success.ToInt(),
		Message: "successful",
		Detail:  nil,
		Data:    data,
	}
}

// 一个id
type SingleId struct {
	Id string `json:"id"`
}

// 一个token
type SingleToken struct {
	Token string `json:"token"`
}
