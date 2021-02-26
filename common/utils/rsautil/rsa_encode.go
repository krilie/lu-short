package rsautil

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"errors"
)

// RsaEncryptOutBase64 对法人信息等敏感字段进行rsa加密
func RsaEncryptOutBase64(str string, key *rsa.PublicKey) (string, error) {
	bytes, e := rsa.EncryptPKCS1v15(rand.Reader, key, []byte(str))
	if e != nil {
		return "", e
	} else {
		return base64.StdEncoding.EncodeToString(bytes), nil
	}
}

// RsaSignStringOutBase64 签名
func RsaSignStringOutBase64(str string, key *rsa.PrivateKey) (string, error) {
	hash := sha1.New()
	hash.Write([]byte(str))
	sign, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA1, hash.Sum(nil))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(sign), nil
}

// RsaCheckSign 验证
func RsaCheckSign(paramStr, base64Sign string, key *rsa.PublicKey) error {
	sign, e := base64.StdEncoding.DecodeString(base64Sign)
	if e != nil {
		return errors.New("base64解码错误:" + e.Error())
	}
	hashd := sha1.Sum([]byte(paramStr))
	e = rsa.VerifyPKCS1v15(key, crypto.SHA1, hashd[:], sign)
	if e != nil {
		return errors.New("rsa签名校验错误:" + e.Error())
	} else {
		return nil
	}
}
