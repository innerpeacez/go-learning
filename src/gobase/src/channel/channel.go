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

}
