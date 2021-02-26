package errs

import (
	"errors"
	"testing"
)

func TestGetErr(t *testing.T) {
	err := NewNormal().WithError(NewNoPermission().WithError(errors.New("ttttt")).WithMsg("internal")).WithMsg("hello")
	t.Log(GetInnerErr(err))
	t.Log(GetErrMsg(err))
	t.Log(errors.Is(err, InternalError))
	t.Log(errors.Is(err, NormalError))
	t.Log(errors.Is(err, NoPermissionError))
	t.Log(errors.As(err, &NormalError))
	var newErr = &Err{}
	t.Log(errors.As(err, &newErr))
}
