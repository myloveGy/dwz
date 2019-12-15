package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_DRIVER   = "mysql"
	DB_USERNAME = "root"
	DB_PASSWORD = "Liujinxing"
	DB_HOST     = "127.0.0.1"
	DB_PORT     = "3306"
	DB_NAME     = "short-url"
	DB_CHARSET  = "utf8"
)

var DB *sql.DB

func init() {

	// 连接数据库
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s",
		DB_USERNAME,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_NAME,
		DB_CHARSET,
	)

	var err error
	DB, err = sql.Open(DB_DRIVER, dsn)
	if err != nil {
		panic(err)
	}
}
