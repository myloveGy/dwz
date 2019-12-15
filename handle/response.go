package handle

import (
	"encoding/json"
	"net/http"
)

// 响应json数据
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type indexResponse struct {
	Date string `json:"date"`
	Time int64  `json:"time"`
	Uri  string `json:"uri"`
}

// 响应错误信息
func responseError(w http.ResponseWriter, code int, msg string) error {
	if msg == "" {
		msg = "fail"
	}

	return responseJson(w, code, msg, nil)
}

func responseSuccess(w http.ResponseWriter, data interface{}, params ...string) error {
	var message string
	if len(params) == 0 || params[0] == "" {
		message = "success"
	} else {
		message = params[0]
	}

	responseJson(w, 200, message, data)
	return nil
}

func responseJson(w http.ResponseWriter, code int, message string, data interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	str, _ := json.Marshal(Response{
		Code: code,
		Msg:  message,
		Data: data,
	})

	w.Write(str)
	return nil
}
