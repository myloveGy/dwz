package handle

import (
	"app/config"
	"app/models"
	"app/utils"
	"net/http"
	"strings"
)

type createResponse struct {
	Url       string `json:"url"`
	ShortUrl  string `json:"short_url"`
	ShortId   string `json:"short_id"`
	CreatedAt string `json:"created_at"`
}

func Create(w http.ResponseWriter, r *http.Request) error {
	r.ParseForm()

	// 获取信息
	key := r.PostForm.Get("app_key")
	url := r.PostForm.Get("url")
	if key == "" || url == "" {
		return responseError(w, 501, "请求参数存在问题")
	}

	// 验证url
	if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://") {
		return responseError(w, 501, "url地址无效")
	}

	// 查询数据是否存在
	app, err := models.FindAppByAppKey(key)
	if err != nil {
		return responseError(w, 502, "没有查询到应用")
	}

	// 验证应用状态
	if app.Status != 1 {
		return responseError(w, 502, "该应用已经被停用")
	}

	// 响应数据
	currentTime := utils.GetCurrentDateTime()
	var appUrl = models.AppUrl{
		AppId:     app.Id,
		Url:       url,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	// 判断是否存在
	var has bool = appUrl.FindOne()

	// 不存在、那么新增数据
	if !has {
		appUrl.Id, err = appUrl.Create()
		if err != nil {
			return err
		}
	}

	// 需要处理ID
	shortId := utils.EncodeInt64(appUrl.Id)
	return responseSuccess(w, createResponse{
		Url:       url,
		ShortUrl:  config.APP_URL + "/" + shortId,
		ShortId:   shortId,
		CreatedAt: appUrl.CreatedAt,
	})
}
