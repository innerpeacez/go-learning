package string

import (
	"testing"
)

func TestString(t *testing.T) {
	var s string

	t.Log(s) // string 是数据类型，零值是""
	// string 是不可变的 byte slice , 使用len 函数获取的是 byte的个数
	s = "hello"
	t.Log(s)
	t.Log(len(s))

	s = "0xhwert"
	t.Log(s)

	t.Log(len(s))

	s = "翟"

	t.Logf("len %d, UTF8 %x", len(s), s)
	c := []rune(s)
	t.Logf("len %d,unicode %x", len(c), c[0])
}
