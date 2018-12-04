package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func main() {
	// 定义int数组切片
	a := []int{1, 2, 3, 4, 5, 6}
	// 创建channel
	c := make(chan int)
	go sum(a, c)
	go sum(a[len(a)/2:], c)
	// 先执行完成的先接受
	sum1, sum2 := <-c, <-c
	fmt.Println(sum1, sum2)

	// 创建缓冲channel

	channel := make(chan int, 3)
	//go sum(a, channel)
	//go sum(a[len(a)/2:], channel)
	//go sum(a[:len(a)/2], channel)
	channel <- 1
	channel <- 2

	// 缓冲区满的时候直接报错
	//channel <- 3

	//for i := range channel {
	//	fmt.Println(i)
	//}

	close(channel)
	v, ok := <-channel
	fmt.Println(ok)
	channel <- 3
	if ok {
		fmt.Println(v)
	}

}
