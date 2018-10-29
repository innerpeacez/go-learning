package main

import "fmt"

func main() {
	pipe = make(chan int, 1)
	go add(2, 5)

	sum := <-pipe
	fmt.Printf("sum = ", sum)
}

var pipe chan int

func add(a int, b int) {
	var sum = a + b
	pipe <- sum
}
