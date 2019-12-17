package models

import (
	"app/config"
	"app/utils"
	"database/sql"
)

type AppUrl struct {
	Id        int64  `db:"id"`
	AppId     int64  `db:"app_id"`
	ShortId   string `db:"short_id"`
	Url       string `db:"url"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

// 新增数据
func (u *AppUrl) Create() (lastId int64, err error) {
	var (
		smt      *sql.Stmt
		result   sql.Result
		querySql = "INSERT INTO `app_url` (`app_id`, `url`, `created_at`, `updated_at`) VALUES (?, ?, ?, ?)"
	)

	smt, err = config.DB.Prepare(querySql)
	if err != nil {
		return
	}

	result, err = smt.Exec(u.AppId, u.Url, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		return
	}

	return result.LastInsertId()
}

func (u *AppUrl) UpdateShortId() (lastId int64, err error) {
	var (
		smt      *sql.Stmt
		result   sql.Result
		querySql = "UPDATE `app_url` SET `short_id` = ? , `updated_at` = ? WHERE `id` = ?"
	)

	smt, err = config.DB.Prepare(querySql)
	if err != nil {
		return
	}

	result, err = smt.Exec(u.ShortId, utils.GetCurrentDateTime(), u.Id)
	if err != nil {
		return
	}

	return result.RowsAffected()
}

// 判断是否存在
func (u *AppUrl) FindOne() bool {
	existsSql := "SELECT `id`, `app_id`, `short_id`, `url`, `created_at`, `updated_at` FROM `app_url` WHERE `app_id` = ? AND `url` = ?"
	_ = config.DB.QueryRow(existsSql, u.AppId, u.Url).Scan(&u.Id, &u.AppId, &u.ShortId, &u.Url, &u.CreatedAt, &u.UpdatedAt)
	return u.Id > 0
}

// 查询数据
func FindUrlById(id int64) (url string, err error) {
	// 查询数据
	querySql := "SELECT `url` FROM `app_url` WHERE `id` = ? LIMIT 1"
	err = config.DB.QueryRow(querySql, id).Scan(&url)
	return
}

// 查询通过short_id
func FindUrlByShortId(shortId string) (url string, err error) {
	querySql := "SELECT `url` FROM `app_url` WHERE `short_id` = ? LIMIT 1"
	err = config.DB.QueryRow(querySql, shortId).Scan(&url)
	return
}
