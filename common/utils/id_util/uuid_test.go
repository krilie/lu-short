package id_util

import (
	uuid "github.com/satori/go.uuid"
	"testing"
)

func TestGetUuid(t *testing.T) {
	t.Log(GetUuid())
	t.Log(GetUuid())
	t.Log(GetUuid())
	t.Log(GetUuid())
	t.Log(uuid.NewV4().String())
	t.Log(uuid.NewV4().String())
	t.Log(uuid.NewV4().String())
}
