package config

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	if dbUser := Get("DB_USERNAME"); dbUser != "root" {
		t.Error("get db_username error", dbUser)
	} else {
		fmt.Println("db_user := ", dbUser)
	}
}

func TestRead(t *testing.T) {
	Read("/usr/local/goproject/my-short-url/.env")
	fmt.Println(Config)
	if Get("DB_PASSWORD") != "Liujinxing" {
		t.Error("读取文件失败")
	}
}

func TestGetValue(t *testing.T) {
	Read("/usr/local/goproject/my-short-url/.env")

	value := GetValue("LINE/int", 10000)
	fmt.Println(value, value.(int))

	value = GetValue("LINE/int64", int64(1))
	fmt.Println(value, value.(int64))
}
