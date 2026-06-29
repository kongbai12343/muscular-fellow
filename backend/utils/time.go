package utils

import "time"

const TimeFormat = "2006-01-02 15:04:05"

var ChinaLocation = time.FixedZone("Asia/Shanghai", 8*60*60)

func FormatDateTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.In(ChinaLocation).Format(TimeFormat)
}
