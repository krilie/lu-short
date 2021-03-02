package model

import (
	"lu-short/common/com_model"
	"time"
)

type TbVisitor struct {
	com_model.TbCommon
	TrackId       string     `json:"track_id" db:"track_id"`
	LastIp        string     `json:"last_ip" db:"last_ip"`
	LastVisitTime *time.Time `json:"last_visit_time" db:"last_visit_time"`
	VisitTimes    int        `json:"visit_times" db:"visit_times"`
}

type TbVisitorLog struct {
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
