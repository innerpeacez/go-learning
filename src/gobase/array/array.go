package main

import "fmt"

func main() {

	var arr [2]int
	arr[1] = 10
	arr[0] = 20

	fmt.Println(arr)

	p := []int{2, 3, 5, 7, 9}

	fmt.Println("p = ", p)
	fmt.Println("p[1:4] = ", p[1:4])
	// 省略下标代表从 0 开始
	fmt.Println("p[:3] ==", p[:3])

	// 省略上标代表到 len(s) 结束
	fmt.Println("p[4:] ==", p[4:])

	ints := make([]int, 5, 1)
	fmt.Println(cap(ints))
	i := ints[cap(ints)]
	fmt.Println(i)

	// 初始化数组
	var arr1 = [2]int{10, 20}
	fmt.Printf("arr1 = ", arr1)

	// 不定义数组的初始容量
	var arr2 = [...]int{10, 20}
	fmt.Printf("arr2 = ", arr2)

}
