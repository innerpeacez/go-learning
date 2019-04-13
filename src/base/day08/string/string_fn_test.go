package string

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {

	s := "A,B,C"
	// string 的分割
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}

	// string 的连接
	join := strings.Join(parts, "/")
	t.Log(join)

}

func TestStr(t *testing.T) {
	// 字符串的转换
	s := "10"
	if v, err := strconv.Atoi(s); err == nil {
		t.Log(v + 10)
	}
	i := 10
	t.Log(strconv.Itoa(i) + "str")

}
