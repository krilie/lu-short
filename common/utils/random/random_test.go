package random

import "testing"

func TestRandomStrWithCharset(t *testing.T) {
	println(GetRandomInt())
	println(GetRandomInt())
	println(GetRandomInt())
	println(GetRandomInt())
	println(GetRandomStr(1))
	println(GetRandomStr(3))
	println(GetRandomStr(9))
	println(GetRandomStr(30))
	println(GetRandomStr(300))
	println(GetRandomStr(300))
	println(GetRandomStr(300))
	println(GetRandomStr(300))
	println(GetRandomStr(300))
	println(GetRandomStrWithCharset(1, "1232123"))
	println(GetRandomStrWithCharset(10, "123456789"))
	println(GetRandomStrWithCharset(100, "abc123"))
	println(GetRandomStrWithCharset(23, "abc"))
	println(GetRandomStrWithCharset(45, "ab"))
	println(GetRandomIntStr())
	println(GetRandomIntStr())
	println(GetRandomIntStr())
	println(GetRandomIntStr())
	println(GetRandomIntStr())
}
