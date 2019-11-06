package main

import "fmt"

func main() {
	add, sub := addAndSub(10, 20)
	fmt.Print(add, sub)
}

func addAndSub(x int, y int) (int, int) {
	var add = x + y
	var sub = x - y
	return add, sub
}
