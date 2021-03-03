package id_util

import (
	"github.com/satori/go.uuid"
	"strings"
)

//没有-的uuid 用做主键
//小写十六进制串 36个 char(36)
func GetUuid() string {
	return strings.ToLower(uuid.NewV4().String())
}
