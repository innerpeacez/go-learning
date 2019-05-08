package main

import "fmt"

func main() {

	i := 100

	// & 符号会生成一个指向其作用对象的指针。也就是底层的内存地址
	p := &i

	fmt.Println(p)
	// * 符号表示指针指向的底层的值。
	fmt.Println(*p)

	*p = 1000
	fmt.Println(*p)
	fmt.Println(i)

	var p2 *int
	p2 = &i
	fmt.Printf("p2 ", p2)
	fmt.Printf("p2 ", *p2)
}
