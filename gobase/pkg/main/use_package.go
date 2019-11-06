package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Print("I Love Go ", rand.Intn(100))

	// 函数名大写表示可以再其他包中访问，类似于java中的public关键字
	fmt.Println(math.Pi)
}
