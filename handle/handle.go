package handle

import (
	"fmt"
	"net/http"
)

func Handle(f func(w http.ResponseWriter, r *http.Request) error) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			// 存在错误、响应处理
			if e := recover(); e != nil {
				fmt.Println("服务器错误:", e)
				responseError(w, 505, "服务器繁忙，请稍后再试...")

			}
		}()

		if err := f(w, r); err != nil {
			responseError(w, 505, err.Error())
		}
	}
}
