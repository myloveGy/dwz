package utils

import (
	"fmt"
	"testing"
	"time"
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

func TestBase62(t *testing.T) {
	str := Base62(1)
	fmt.Println(str, Base62(2), Base62(23), Base62(32), Base62(63), Base62(time.Now().Unix()))
}
