package handle

import "time"

const OneMinute = 60
const OneHour = 60 * OneMinute
const OneDay = 24 * OneHour

var GetNowTime = func() *time.Time {
	t := time.Now()
	t = TranslateTime(t)
	return &t
}

var TranslateTime = func(t time.Time) time.Time {
	value, _ := time.LoadLocation("Asia/Shanghai")
	return t.In(value)
}

func ParseTime(unix int64) *time.Time {
	if unix == 0 {
		return nil
	}
	t := time.UnixMilli(unix)
	return &t
}

func ToUnix(t *time.Time) int64 {
	if t == nil {
		return 0
	}
	return t.UnixMilli()
}
