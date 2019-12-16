package main

import (
	"app/config"
	"app/handle"
	"fmt"
	"net/http"
)

func main() {
	// 读取配置文件，并且配置DB
	config.Read("./.env")
	config.RegisterDB()

	// 配置路由
	http.HandleFunc("/", handle.Handle(handle.Home))
	http.HandleFunc("/create", handle.Handle(handle.Create))

	// 开启Http服务
	fmt.Println("Listen Port:", config.Get("APP_PORT"))
	http.ListenAndServe(config.Get("APP_PORT", ":80"), nil)
}
