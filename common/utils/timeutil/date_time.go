package timeutil

import (
	"gorm.io/gorm"
	"time"
)

// 本月开始时间
func GetBeijingMonthStartTime(time time.Time) time.Time {
	time = time.In(BeijingZone).AddDate(0, 0, -time.Day()+1)
	return GetBeijingZeroTime(time)
}

// 获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetBeijingLastDateOfMonth(d time.Time) time.Time {
	return GetBeijingMonthStartTime(d).In(BeijingZone).AddDate(0, 1, 0).Add(-time.Second)
}

// 获取零点时间
func GetBeijingZeroTime(d time.Time) time.Time {
	d = d.In(BeijingZone)
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

func SqlNullTime(t *time.Time) gorm.DeletedAt {
	if t != nil {
		return gorm.DeletedAt{
			Time:  *t,
			Valid: true,
		}
	}
	return gorm.DeletedAt{
		Time:  time.Time{},
		Valid: false,
	}
}
