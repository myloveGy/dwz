package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func RegisterDB() {

	// 连接数据库
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s",
		Get("DB_USERNAME", "root"),
		Get("DB_PASSWORD"),
		Get("DB_HOST"),
		Get("DB_PORT"),
		Get("DB_NAME"),
		Get("DB_CHARSET"),
	)

	var err error
	DB, err = sql.Open(Get("DB_DRIVER", "mysql"), dsn)
	if err != nil {
		panic(err)
	}
}
