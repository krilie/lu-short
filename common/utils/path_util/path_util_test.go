package path_util

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAutoGetWd(t *testing.T) {
	t.Log(GetRootWd(0))
	t.Log(GetRootWd(1))
	t.Log(GetRootWd(2))
	t.Log(GetRootWd(3)) // 项目目录
	t.Log(GetRootWd(4)) // 项目根目录
	require.Panics(t, func() {
		t.Log(GetRootWd(6))
	})
}

func TestAuto_panicOnErr(t *testing.T) {
	require.Panics(t, func() {
		panicOnErr(errors.New("panic test"))
	})
	require.NotPanics(t, func() {
		panicOnErr(nil)
	})
}

func TestAutoGetRootWdBefore(t *testing.T) {
	t.Log(GetRootWdBefore("common"))
	require.Panics(t, func() {
		t.Log(GetRootWdBefore("common2"))
	})
	t.Log(GetRootWdBefore("utils"))
	t.Log(GetRootWdBefore("path_util"))
}
