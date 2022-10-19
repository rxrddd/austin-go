package timex

import (
	"github.com/spf13/cast"
	"time"
)

const DateTimeLayout = "2006-01-02 15:04:05"
const DateLayout = "2006-01-02"

func NowDateTime() string {
	return time.Now().Format(DateTimeLayout)
}
func NowDate() string {
	return time.Now().Format(DateLayout)
}
func FormatDate(i interface{}) string {
	switch v := i.(type) {
	case time.Time:
		if v.IsZero() {
			return ""
		}
		return v.Format(DateLayout)
	case *time.Time:
		if v != nil {
			if v.IsZero() {
				return ""
			}
			return v.Format(DateLayout)
		}
	}
	return ""
}

func FormatDateTime(i interface{}) string {
	switch v := i.(type) {
	case time.Time:
		if v.IsZero() {
			return ""
		}
		return v.Format(DateTimeLayout)
	case *time.Time:
		if v != nil {
			if v.IsZero() {
				return ""
			}
			return v.Format(DateTimeLayout)
		}
	}
	return ""
}

func Parse(str string) time.Time {
	parse, err := time.Parse(DateTimeLayout, cast.ToString(str))
	if err != nil {
		panic(err)
	}
	return parse
}

// GetDisTodayEnd 获取现在到今天结束的秒数
func GetDisTodayEnd() int64 {
	todayEnd, _ := time.ParseInLocation("2006-01-02", time.Now().Format("2006-01-02"), time.Local)
	todayEndUnix := todayEnd.AddDate(0, 0, 1).Unix()
	period := todayEndUnix - time.Now().Unix()
	return period
}
