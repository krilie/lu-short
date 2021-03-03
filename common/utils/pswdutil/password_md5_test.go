package pswdutil

import (
	"fmt"
	"testing"
)

func TestGetMd5Password(t *testing.T) {
	t.Log(GetMd5Password("12345678", "234343"))
	t.Log(GetMd5Password("12345678", "234343"))
	t.Log(GetMd5Password("12345678", "234342"))
}

func TestIsPasswordOk(t *testing.T) {
	password := GetMd5Password("123", "123")
	isOk := IsPasswordOk("123", password, "123")
	if !isOk {
		t.Error("bad on check md5 pswd:", "123", "123", password)
	} else {
		t.Log("md5 checked ok", password)
	}
}

func TestGetSalt(t *testing.T) {
	fmt.Println(GetSalt(1))
	fmt.Println(GetSalt(10))
	fmt.Println(GetSalt(4))
	fmt.Println(GetSalt(18))
	fmt.Println(GetSalt(23))
	fmt.Println(GetSalt(7))
}
