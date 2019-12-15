package models

import "app/config"

type App struct {
	Id        int64  `db:"app_id"`
	AppName   string `db:"app_name"`
	AppKey    string `db:"app_key"`
	Status    int    `db:"status"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

func FindAppByAppKey(appKey string) (App, error) {
	var (
		app App
		err error
	)

	// 查询数据
	err = config.DB.QueryRow("SELECT * FROM `app` WHERE `app_key` = ? LIMIT 1", appKey).Scan(
		&app.Id,
		&app.AppName,
		&app.AppKey,
		&app.Status,
		&app.CreatedAt,
		&app.UpdatedAt,
	)

	return app, err
}
