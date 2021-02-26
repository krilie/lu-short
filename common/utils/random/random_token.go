package random

//给app用户用的random_token生成
//要求短且好看

//获取盐值
func GetRandomNum(size int) string {
	if size == 0 {
		return ""
	}
	return GetRandomStrWithCharset(1, "123456789") + GetRandomStrWithCharset(size-1, "1234567890")
}
