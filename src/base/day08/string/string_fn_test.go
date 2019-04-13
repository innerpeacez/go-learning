package string

import (
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {

	s := "A,B,C"
	// 字符串的分割
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}

}
