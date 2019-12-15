package utils

import "time"

const (
	FORMAT_DATETIME = "2006-01-02 15:04:05"
	FORMAT_DATE     = "2006-01-02"
)

func getFirstValue(defaultValue string, format ...string) string {
	if len(format) > 0 && format[0] != "" {
		defaultValue = format[0]
	}

	return defaultValue
}

// 获取当前日期时间
func GetCurrentDateTime(format ...string) string {
	return time.Now().Format(getFirstValue(FORMAT_DATETIME))
}

// 获取当前日期
func GetCurrentDate(format ...string) string {
	return time.Now().Format(getFirstValue(FORMAT_DATE))
}
