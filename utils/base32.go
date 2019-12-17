package utils

import (
	"encoding/base32"
	"math/rand"
	"strconv"
	"time"
)

const BASE = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

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

func Base62(id int64) string {
	if id <= 0 {
		return ""
	}

	var str = []byte{}
	for id > 0 {
		str = append(str, BASE[id%62])
		id = id / 62
	}

	i := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(51)
	str = append(str, BASE[i+10])
	return string(str)
}
