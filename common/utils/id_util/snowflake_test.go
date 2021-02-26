package id_util

import (
	"errors"
	"fmt"
	"testing"
)

func TestNextSnowflakeId(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(NextSnowflakeId().String())
	}
}

type Err struct {
	msg string
}

func (Err) Error() string {
	return "ok"
}

func TestNewDao(t *testing.T) {
	err := &Err{msg: "l"}
	is := errors.As(err, new(*Err))
	t.Log(is)
}
