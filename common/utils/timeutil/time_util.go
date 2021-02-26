package timeutil

import "time"

var BeijingZone *time.Location

var DefaultFormat = "2006-01-02 15:04:05"
var StringFormat = "20060102150405"

func init() {
	// 东八区北京时间
	BeijingZone = time.FixedZone("CST", 8*3600)
}

func GetBeijingNowTimeString() string {
	return time.Now().In(BeijingZone).Format("2006-01-02 15:04:05")
}
func GetBeijingNowTimeStringFormat(format string) string {
	return time.Now().In(BeijingZone).Format(format)
}

func GetTimeNow() *time.Time {
	timeN := time.Now().In(BeijingZone)
	return &timeN
}

func GetBeijingTimeString(unix int64) string {
	return time.Unix(unix, 0).In(BeijingZone).Format("2006-01-02 15:04:05")
}
func GetNowUtcTimeString() string {
	return time.Now().In(time.UTC).Format("2006-01-02 15:04:05")
}

func ParseBeijingTime(str, format string) (time.Time, error) {
	return time.ParseInLocation(format, str, BeijingZone)
}
