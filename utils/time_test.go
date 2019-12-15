package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestGetCurrentDateTime(t *testing.T) {
	str := GetCurrentDateTime()
	if str != time.Now().Format(FORMAT_DATETIME) {
		t.Error("获取日期时间错误", str)
	}

	fmt.Println("当前日期时间为:", str)
}

func TestGetCurrentDate(t *testing.T) {
	str := GetCurrentDate()
	if str != time.Now().Format(FORMAT_DATE) {
		t.Error("获取日期错误:", str)
	}

	fmt.Println("当前日期为:", str)
}

func TestGetFirstValue(t *testing.T) {
	str := getFirstValue("123", "456")
	if str != "456" {
		t.Error("获取值错误:", str)
	}

	fmt.Println("获取的值为:", str, "没有使用默认值")

	str = getFirstValue("789", "", "456")
	if str != "789" {
		t.Error("获取值错误:", str)
	}

	fmt.Println("获取的值为:", str, "使用默认值")

	str = getFirstValue("789")
	if str != "789" {
		t.Error("获取值错误:", str)
	}

	fmt.Println("获取的值为:", str, "使用默认值")
}
