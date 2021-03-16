package ginutil

import (
	"lu-short/common/com_model"
	"lu-short/common/errs"
)

// 处理错误，如果有错误返回真 无错误返回假
func (g *GinWrap) HandlerError(err error) bool {
	if err == nil {
		return false
	} else {
		g.ReturnWithErr(err)
		return true
	}
}

// 处理错误 如果没有返回通用成功
func (g *GinWrap) HandlerErrorOrReturnSuccess(err error) {
	if err == nil {
		g.Context.JSON(200, com_model.StdSuccess)
		return
	} else {
		g.ReturnWithErr(err)
		return
	}
}

// 处理错误 如果没有返回通用成功
func (g *GinWrap) HandlerErrorOrReturnJson(err error, ret interface{}) {
	if err == nil {
		g.Context.JSON(200, ret)
		return
	} else {
		g.ReturnWithErr(err)
		return
	}
}

// 处理错误 如果没有返回通用成功
func (g *GinWrap) HandlerErrorOrReturnData(err error, data interface{}) {
	if err == nil {
		g.ReturnData(data)
		return
	} else {
		g.ReturnWithErr(err)
		return
	}
}

// abort with err use err's default http status
func (g *GinWrap) ReturnWithErr(err error) {
	if nErr := errs.ToErrOrNil(err); nErr != nil {
		g.Context.JSON(200, com_model.NewRet(nErr))
	} else {
		g.Context.JSON(200, com_model.NewRetFromErr(err))
	}
}

func (g *GinWrap) ReturnWithAppErr(err *errs.Err) {
	g.Context.JSON(200, com_model.NewRet(err))
}

func (g *GinWrap) ReturnWithParamErr(err error) {
	g.Context.JSON(200, com_model.NewFailure(errs.ErrorParam, err.Error()))
}

func (g *GinWrap) ReturnOk() {
	g.Context.JSON(200, com_model.StdSuccess)
}

func (g *GinWrap) ReturnData(data interface{}) {
	g.Context.JSON(200, com_model.NewSuccess(data))
}

func (g *GinWrap) ReturnFailure(code errs.ErrCode, msg string) {
	g.Context.JSON(200, com_model.NewFailure(code, msg))
}
