package utils

import (
	"fmt"
	"testing"
)

func TestDecodeInt64(t *testing.T) {
	id, err := DecodeInt64("GI======")
	if err != nil {
		t.Error("解析错误", err)
		return
	}

	if id != 2 {
		t.Error("解析失败:", id)
		return
	}

	fmt.Println(id)
}

func TestEncodeInt64(t *testing.T) {
	str := EncodeInt64(2)
	if str != "GI======" {
		t.Error("加密失败", str)
		return
	}

	fmt.Println(str)
}
