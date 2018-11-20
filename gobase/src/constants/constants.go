package main

import "fmt"

const PI = 3.141592653

const (
	Name = "innerpeacez"
	Age  = 14
)

func main() {
	const HelloWorld = "你好,世界"
	// PI = 1234,常量无法修改左值
	fmt.Println(PI)
	fmt.Println(HelloWorld)
}
