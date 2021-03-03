package jwt

import (
	"errors"
	"time"
)

//同一類型的多個錯誤實例
var ErrTimeExp = errors.New("jwt is expiration")
var ErrIatTime = errors.New("jwt in bad format,iat>=exp")

//jwt payload的內容，如果是app角色，則appid為空。,web端默認web
type UserClaims struct {
	ClientId string `json:"client_id"` //頒發給哪個端
	UserId   string `json:"user_id"`   //用戶id
	Iat      int64  `json:"iat"`       //發放時間
	Exp      int64  `json:"exp"`       //過期時間
	Jti      string `json:"jti"`       //token 的id ,唯一標識
	Iss      string `json:"iss"`       //簽發者 是sys_user_control
}

//jwt 是否有效，如果沒效，則入出錯誤
func (u UserClaims) Valid() error {
	if time.Now().Unix() > u.Exp {
		return ErrTimeExp
	}
	if u.Iat > u.Exp {
		return ErrIatTime
	}
	return nil
}
