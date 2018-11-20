package main

import "fmt"

const PI = 3.141592653

func main() {
	const HELLO_WORLD = "你好，世界"
	// PI = 1234,常量无法修改左值
	fmt.Println(PI)
	fmt.Println(HELLO_WORLD)
}
