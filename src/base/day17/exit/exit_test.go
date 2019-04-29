package exit

import (
	"os"
	"testing"
)

// os.Exit 函数停止时不会执行 defer 函数，并且不会打印出调用栈
func TestExit(t *testing.T) {
	t.Log("start")

	defer func() {
		t.Log("defer func")
	}()

	os.Exit(-1)
}
