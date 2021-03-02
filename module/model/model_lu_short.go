package model

import (
	"lu-short/common/com_model"
	"time"
)

type TbRedirect struct {
	com_model.TbCommon `json:"" json:""`
	CustomerId         string    `json:"customer_id" json:"customer_id"`
	OriUrl             string    `json:"ori_url" json:"ori_url"`
	Key                string    `json:"key" json:"key"`
	RateLimit          int       `json:"rate_limit" json:"rate_limit"`             // 每秒访问次数
	TimesLimitLeft     int       `json:"times_limit_left" json:"times_limit_left"` // 总访问次数限制
	JumpLimitLeft      int       `json:"jump_limit_left" json:"jump_limit_left"`   // 跳转次数限制
	BeginTime          time.Time `json:"begin_time" json:"begin_time"`             // 有效开始时间
	DeadTime           time.Time `json:"dead_time" json:"dead_time"`               // 有效时间
}
