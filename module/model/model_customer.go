package model

import "lu-short/common/com_model"

type TbCustomer struct {
	com_model.TbCommon
	LoginName string `json:"login_name" db:"login_name"`
	PhoneNum  string `json:"phone_num" db:"phone_num"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	Picture   string `json:"picture" db:"picture"`
	Salt      string `json:"salt" db:"salt"`
	Vip       int    `json:"vip" db:"vip"`
}
