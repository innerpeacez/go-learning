package inteface

import (
	"fmt"
	"testing"
)

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
	programmer := &GoProgrammer{}
	hello := programmer.WriteHelloWorld()
	t.Log(hello)
}

func hello() {
	fmt.Println("hello")
}

// 自定义接口类型
func Hello(hello func()) string {
	hello()
	return "hello"
}

type hello2 func()

func (hello2 *hello2) hhello2() {
	fmt.Println("hello 2")
}

func Hello2(hello2 hello2) string {
	hello2()
	return "hello"
}

func TestDefinition(t *testing.T) {
	Hello(hello)

}
