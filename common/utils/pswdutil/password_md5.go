package pswdutil

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/satori/go.uuid"
	"strings"
)

//对密码进行md5加密 盐值
//原密码加盐值，返回加密过后的密码,大写十六进制串
func GetMd5Password(ori string, salt string) string {
	tagPswd := md5.Sum([]byte(ori + salt + ori + "#*" + salt))
	tagPswdHex := strings.ToLower(hex.EncodeToString(tagPswd[:]))
	return tagPswdHex
}

//检查password是否正常
func IsPasswordOk(ori, md5ed, salt string) bool {
	return GetMd5Password(ori, salt) == md5ed
}

func Md5(str string) string {
	tagPswd := md5.Sum([]byte(str))
	tagPswdHex := strings.ToLower(hex.EncodeToString(tagPswd[:]))
	return tagPswdHex
}

//获取盐值
func GetSalt(num uint) string {
	randomStrings := hex.EncodeToString(uuid.NewV4().Bytes())
	return randomStrings[0:num]
}
