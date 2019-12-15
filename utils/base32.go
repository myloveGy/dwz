package utils

import (
	"encoding/base32"
	"strconv"
)

// 将字符串解析为int64
func DecodeInt64(uri string) (id int64, err error) {
	var strId []byte
	strId, err = base32.StdEncoding.DecodeString(uri)
	if err != nil {
		return
	}

	return strconv.ParseInt(string(strId), 10, 64)
}

// 将int64转为string
func EncodeInt64(id int64) string {
	str := strconv.FormatInt(id, 10)
	return base32.StdEncoding.EncodeToString([]byte(str))
}
