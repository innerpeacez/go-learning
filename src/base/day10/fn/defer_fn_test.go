package fn

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	// defer 函数
	defer func() {
		t.Log("Clear resources")
	}()

	t.Log("Started")
	//panic("Fatal error") // 抛出异常
}

func Clear() {
	fmt.Println("Clear resources")
}

func TestDefer2(t *testing.T) {
	defer Clear() // panic 时defer函数也会执行
	fmt.Println("Start")
	//panic("error")
}
