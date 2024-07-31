package utils

import (
	"time"
)

var cstZone = time.FixedZone("CST", 8*3600)

const (
	date             = "20060102"
	dateMs           = "20060102150405"
	timeTemplateDate = "2006-01-02"
	timeTemplateTime = "2006-01-02 15:04:05"
	OneWeekMills     = 7 * 24 * 60 * 60 * 1000
)

func Microsecond() int64 {
	return time.Now().UnixNano() / 1e6
}

func Day(timestamp int64) int64 {
	t := time.Unix(timestamp/1000, 0)
	day := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, cstZone)
	return day.UnixNano() / 1e6
}

func Today() int64 {
	t := time.Now()
	day := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, cstZone)
	return day.UnixNano() / 1e6
}

func YesterdayStart() int64 {
	t := time.Now()
	day := time.Date(t.Year(), t.Month(), t.Day()-1, 0, 0, 0, 0, cstZone)
	return day.UnixNano() / 1e6
}
func YesterdayEnd() int64 {
	t := time.Now().In(cstZone)
	day := time.Date(t.Year(), t.Month(), t.Day()-1, 23, 59, 59, 59, cstZone)
	return day.UnixNano() / 1e6
}

func TomorrowStart(timestamp int64) int64 {
	t := time.Unix(timestamp/1e3, 0)
	day := time.Date(t.Year(), t.Month(), t.Day()+1, 0, 0, 0, 0, cstZone)
	return day.UnixNano() / 1e6
}

func ThisMonthStart() int64 {
	t := time.Now()
	day := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, cstZone)
	return day.UnixNano() / 1e6
}
func ThisYearStart() int64 {
	t := time.Now()
	day := time.Date(t.Year(), 1, 1, 0, 0, 0, 0, cstZone)
	return day.UnixNano() / 1e6
}
func GetHourStart(t time.Time) int64 {
	day := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, cstZone)
	return day.UnixNano() / 1e6
}

func GetDayStart(t time.Time) int64 {
	day := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, cstZone)
	return day.UnixNano() / 1e6
}

func GetDay() string {
	return time.Now().In(cstZone).Format(date)
}

func FormatDate(t int64) string {
	return time.Unix(t/1e3, 0).Format(timeTemplateDate)
}

func FormatTime(t int64) string {
	return time.Unix(t/1e3, 0).Format(timeTemplateTime)
}

func ParseDate(date string) (int64, error) {
	t, err := time.ParseInLocation(timeTemplateDate, date, cstZone)
	if err != nil {
		return 0, err
	}
	return t.UnixNano() / 1e6, nil
}
func GetCurrentTimeFormat() string {
	return time.Now().Format("2006/01/02 15:04")
}

func ThisWeekStart() int64 {
	t := time.Now()
	week := int(t.Weekday() - 1)
	day := time.Date(t.Year(), t.Month(), t.Day()-week, 0, 0, 0, 0, cstZone)
	return day.UnixNano() / 1e6
}

func GetDayMs() string {
	return time.Now().In(cstZone).Format(dateMs)
}

func TimeStr(timeString, desiredFormat string) (string, error) {
	parsedTime, err := time.Parse(time.RFC3339, timeString)
	if err != nil {
		return "", err
	}
	formattedTime := parsedTime.Format(desiredFormat)

	return formattedTime, nil
}

func ConvertToUTC(timeStr string, layout string) (*time.Time, error) {
	// 解析时间字符串
	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		return &time.Time{}, err
	}

	// 加载 "Etc/GMT" 时区
	etcGMT, err := time.LoadLocation("Etc/GMT")
	if err != nil {
		return &time.Time{}, err
	}

	// 将时间转换为 "Etc/GMT" 时区
	etcGMTTime := parsedTime.In(etcGMT)

	// 将时间转换为 UTC 时区
	utcTime := etcGMTTime.UTC()

	return &utcTime, nil
}

func FormatDateToYMD(dateStr string) string {
	// 解析 ISO 8601 日期字符串
	t, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return ""
	}
	return t.Format("2006-01-02")
}

// ConvertTimeFormat 将 ISO 8601 格式的时间字符串转换为指定的格式
func ConvertTimeFormat(input string) (string, error) {
	// 解析时间字符串
	t, err := time.Parse(time.RFC3339, input)
	if err != nil {
		return "", err
	}

	// 格式化时间
	output := t.Format("01/02/06")
	return output, nil
}

// IsBeforeCurrentDate 判断给定时间是否晚于当前时间（只比较到天）
func IsBeforeCurrentDate(targetTime time.Time) bool {
	// 获取当前时间
	currentTime := time.Now()

	// 只取日期部分，忽略时、分、秒
	currentTime = time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
	targetTime = time.Date(targetTime.Year(), targetTime.Month(), targetTime.Day(), 0, 0, 0, 0, targetTime.Location())

	// 比较时间
	return targetTime.Before(currentTime)
}

// ConvertToDateOnly 将日期时间字符串转换为仅包含年月日的字符串
func ConvertToDateOnly(dateTimeStr string) (string, error) {
	// 定义日期时间格式
	layout := time.RFC3339

	// 解析日期时间字符串
	parsedTime, err := time.Parse(layout, dateTimeStr)
	if err != nil {
		return "", err
	}

	// 格式化为仅包含年月日的字符串
	dateOnly := parsedTime.Format("2006-01-02")

	return dateOnly, nil
}
