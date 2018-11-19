package main

import "fmt"

var num1, num2 = 1, 2

func main() {
	// 使用var 定义变量时可以省略类型
	var c, py, java = true, false, "good"
	fmt.Print(c, py, java)
}
