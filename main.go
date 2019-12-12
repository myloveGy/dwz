package main

import (
	"encoding/base32"
	"encoding/json"
	"net/http"
	"strings"
	"time"
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
func responseError(w http.ResponseWriter, code int, msg string) {
	if msg == "" {
		msg = "fail"
	}

	responseJson(w, code, msg, nil)
}

func responseSuccess(w http.ResponseWriter, data interface{}, msg string) {
	if msg == "" {
		msg = "success"
	}

	responseJson(w, 200, msg, data)
}

func responseJson(w http.ResponseWriter, code int, message string, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	str, _ := json.Marshal(Response{
		Code: code,
		Msg:  message,
		Data: data,
	})

	w.Write(str)
}

func Index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// 获取请求地址
	uri := strings.TrimLeft(r.RequestURI, "/")
	if uri == "" {
		responseError(w, 504, "没有对应的跳转地址")
		return
	}

	// 解析ID
	id, err := base32.StdEncoding.DecodeString(uri)
	if err != nil {
		responseError(w, 505, err.Error())
		return
	}

	// 拿到ID查询数据库
	w.Write(id)
}

func Create(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// 获取信息
	key := r.PostForm.Get("app_key")
	url := r.PostForm.Get("url")
	if key == "" || url == "" {
		responseError(w, 501, "请求参数存在问题")
		return
	}

	// 验证秘钥
	if key != "1234567890" {
		responseError(w, 502, "应用Key错误")
		return
	}

	// 响应数据
	currentTime := time.Now()

	// 响应正常信息
	responseSuccess(w, indexResponse{
		Date: currentTime.Format("2006-01-02 15:04:05"),
		Time: currentTime.Unix(),
		Uri:  r.RequestURI,
	}, "")
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/create", Create)
	http.ListenAndServe(":8080", nil)
}
