package model

import (
	"lu-short/common/com_model"
	"time"
)

type TbRedirect struct {
	com_model.TbCommon
	CustomerId     string    `json:"customer_id" db:"customer_id"`
	OriUrl         string    `json:"ori_url" db:"ori_url"`
	Key            string    `json:"key" db:"key"`
	RateLimit      int       `json:"rate_limit" db:"rate_limit"`             // 每秒访问次数
	TimesLimitLeft int       `json:"times_limit_left" db:"times_limit_left"` // 总访问次数限制
	JumpLimitLeft  int       `json:"jump_limit_left" db:"jump_limit_left"`   // 跳转次数限制
	BeginTime      time.Time `json:"begin_time" db:"begin_time"`             // 有效开始时间
	EndTime        time.Time `json:"end_time" db:"end_time"`                 // 有效时间
}

type TbRedirectLog struct {
	com_model.TbCommon
	TrackId    string    `json:"track_id" db:"track_id"`
	Ip         string    `json:"ip" db:"ip"`
	VisitTime  time.Time `json:"visit_time" db:"visit_time"`
	Device     string    `json:"device" db:"device"`
	CustomerId string    `json:"customer_id" db:"customer_id"`
	RedirectId string    `json:"redirect_id" db:"redirect_id"`
	ShortUrl   string    `json:"short_url" db:"short_url"`
	OriUrl     string    `json:"ori_url" db:"ori_url"`
}
