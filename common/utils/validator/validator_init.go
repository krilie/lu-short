package validator

import "github.com/asaskevich/govalidator"

var (
	//密码
	PatternPassword = "[[:alnum:]~!@#$%^&*()-=]{8,50}"
	//登录名
	PatternLoginName = "^[a-zA-Z]{1}[:alnum:]{0,49}"
	//手机号
	PatternPhoneNum = "^(\\+\\d{2}-)?(\\d{2,3}-)?([1][3,4,5,7,8][0-9]\\d{8})$"
	//id号  字母数字32位
	PatternId = "[:alnum:]{32}"
)

func IsPassword(pswd string) bool {
	return govalidator.Matches(pswd, PatternPassword)
}
func IsLoginName(name string) bool {
	return govalidator.Matches(name, PatternLoginName)
}

func IsPhoneNum(num string) bool {
	return govalidator.Matches(num, PatternPhoneNum)
}

func IsIdStr(id string) bool {
	return govalidator.Matches(id, PatternId)
}
