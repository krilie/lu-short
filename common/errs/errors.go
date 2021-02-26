package errs

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type ErrCode int

func (code ErrCode) ToInt() int    { return int(code) }
func (code ErrCode) ToStr() string { return strconv.Itoa(int(code)) }

const (
	Success           ErrCode = 2000
	ErrorNormal       ErrCode = 2500
	ErrorParam        ErrCode = 4000
	ErrorNoPermission ErrCode = 4001
	ErrorInvalidToken ErrCode = 4002
	ErrorNotExists    ErrCode = 4004
	ErrorInternal     ErrCode = 5000
)

// 定义通用错误 pkg errors cause error
var (
	NormalError       = &Err{Code: ErrorNormal, Message: "业务错误"}       // 2100
	ParamError        = &Err{Code: ErrorParam, Message: "参数错误"}        // 4000
	NoPermissionError = &Err{Code: ErrorNoPermission, Message: "无权限"}  // 4001
	InternalError     = &Err{Code: ErrorInternal, Message: "内部错误"}     // 5000
	InvalidTokenError = &Err{Code: ErrorInvalidToken, Message: "凭证无效"} // invalidToken
	NotExistsError    = &Err{Code: ErrorNotExists, Message: "未找到"}     // not found
)

func New(code ErrCode) *Err   { return &Err{Code: code} }
func NewNoPermission() *Err   { return NoPermissionError.New() }
func NewInternal() *Err       { return InternalError.New() }
func NewNormal() *Err         { return NormalError.New() }
func NewInvalidToken() *Err   { return InvalidTokenError.New() }
func NewParamError() *Err     { return ParamError.New() }
func NewNotExistsError() *Err { return NotExistsError.New() }

type Err struct {
	Code      ErrCode
	Message   string
	InsideErr error // 原始错误
}

func (w *Err) Is(target error) bool {
	val, ok := target.(*Err)
	if !ok {
		return false
	}
	return w.Code == val.Code
}

func (w *Err) Error() string {
	builder := strings.Builder{}
	builder.WriteString("[code:")
	builder.WriteString(w.Code.ToStr())
	builder.WriteString(" message:")
	builder.WriteString(w.Message + "]")
	return builder.String()
}

func (w *Err) WithError(err error) *Err   { w.InsideErr = err; return w }
func (w *Err) GetCode() int               { return w.Code.ToInt() }
func (w *Err) WithCode(code ErrCode) *Err { w.Code = code; return w }
func (w *Err) WithMsg(msg string) *Err    { w.Message = msg; return w }
func (w *Err) New() *Err                  { return &Err{Code: w.Code, Message: w.Message, InsideErr: w.InsideErr} }
func (w *Err) Unwrap() error              { return w.InsideErr }

func (w *Err) WithMsgf(format string, args ...interface{}) *Err {
	w.Message = fmt.Sprintf(format, args...)
	return w
}

func (w *Err) GetFullMsg() string {
	return GetErrMsg(w)
}

// 取到最内层的Err 如是没有返回nil
func GetInnerErr(err error) *Err {
	var retErr *Err
	for err != nil {
		tErr, ok := err.(*Err)
		if ok {
			retErr = tErr
			err = tErr.InsideErr
			continue
		} else {
			break
		}
	}
	return retErr
}

func ToErrOrNil(err error) *Err {
	if err == nil {
		return nil
	}
	tErr, ok := err.(*Err)
	if ok {
		return tErr
	}
	return nil
}

func GetCode(err error) int {
	Err := ToErrOrNil(err)
	if Err != nil {
		return Err.Code.ToInt()
	}
	return 5000
}

// 循环取出所有错误信息
func GetErrMsg(err error) string {
	if err == nil {
		return ""
	}
	var msg strings.Builder
	for {
		if err != nil {
			msg.WriteString(err.Error() + "-")
			err = errors.Unwrap(err)
		} else {
			break
		}
	}
	return msg.String()[:msg.Len()-1]
}
