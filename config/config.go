package config

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// 所有配置信息
var Config = map[string]string{
	// 应用配置
	"APP_URL":  "http://localhost",
	"APP_PORT": ":80",

	// 数据库配置
	"DB_DRIVER":   "mysql",
	"DB_HOST":     "127.0.0.1",
	"DB_PORT":     "3306",
	"DB_USERNAME": "root",
	"DB_PASSWORD": "",
	"DB_CHARSET":  "utf8",
	"DB_NAME":     "short-url",
}

func Read(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("读取配置文件错误", err.Error())
		return
	}

	defer file.Close()
	r := bufio.NewReader(file)
	for {
		b, _, err := r.ReadLine()

		// 读取一行存在错误
		if err != nil {
			if err == io.EOF {
				break
			}

			return
		}

		// 去掉两边的空格
		str := strings.TrimSpace(string(b))
		if str == "" {
			continue
		}

		// 查询是否存在=
		index := strings.Index(str, "=")
		if index < 0 {
			continue
		}

		// 存在KEY
		key := strings.TrimSpace(str[:index])
		value := strings.TrimSpace(str[index+1:])
		if len(key) == 0 || len(value) == 0 {
			continue
		}

		// 设置值
		Config[key] = value
	}
}

// 获取配置项
func Get(key string, defaultValue ...string) string {
	str, ok := Config[key]
	// 存在不为空的值，直接返回
	if ok && str != "" {
		return str
	}

	if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return ""
}

func GetValue(key string, defaultValue ...interface{}) interface{} {
	arr := strings.Split(key, "/")
	key = arr[0]
	var typeName string
	if len(arr) > 1 {
		typeName = arr[1]
	} else {
		typeName = "string"
	}

	// 获取到值
	value := Get(key)

	// 转int
	if typeName == "int" {
		intValue, err := strconv.Atoi(value)
		if (intValue == 0 || err != nil) && len(defaultValue) > 0 {
			intValue, _ = defaultValue[0].(int)
		}

		return intValue
	}

	// 转int64
	if typeName == "int64" {
		int64Value, err := strconv.ParseInt(value, 10, 64)
		if (int64Value == 0 || err != nil) && len(defaultValue) > 0 {
			int64Value, _ = defaultValue[0].(int64)
		}

		return int64Value
	}

	if value == "" && len(defaultValue) > 0 {
		value, _ = defaultValue[0].(string)
	}

	return value
}
