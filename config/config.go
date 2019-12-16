package config

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

	return defaultValue[0]
}
