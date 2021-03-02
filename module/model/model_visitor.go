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
