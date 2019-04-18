package inteface

import "testing"

// 定义一个接口
type Programmer interface {
	WriteHelloWorld() string
}

// 初始化一个实例
type GoProgrammer struct {
}

// 实现接口的方法
func (goProgrammer *GoProgrammer) WriteHelloWorld() string {
	hello := "hello"
	return hello
}

// 测试类
func TestInterface(t *testing.T) {
	programmer := GoProgrammer{}
	hello := programmer.WriteHelloWorld()
	t.Log(hello)
}
