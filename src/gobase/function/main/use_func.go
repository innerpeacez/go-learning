package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

// 相同参数类型可以省略
func sub(x, y int) int {
	return x - y
}

func main() {
	fmt.Println(add(10, 20))

	fmt.Println(sub(100, 20))
}
