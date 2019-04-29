package panic

import (
	"errors"
	"testing"
)

// panic 函数停止， 会执行defer 函数，同时会打印出调用栈
func TestPanic(t *testing.T) {
	t.Log("start")

	// 使用 defer 和 recover 模拟 try ... cache
	defer func() {
		t.Log("defer func")

		// recover 函数相当于 java 中的 try ... cache
		if err := recover(); err != nil {
			t.Log("recovered from ", err)
		}
	}()

	panic(errors.New("something fail"))
}
