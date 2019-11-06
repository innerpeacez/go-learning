package main

import "fmt"

func main() {
	sum, avg := calc(100, 200)
	fmt.Printf("sum", sum)
	fmt.Printf("avg", avg)
}

// 两个以上的返回值需要使用()
func calc(a int, b int) (int, int) {
	c := a + b
	d := (a + b) / 2
	return c, d
}
