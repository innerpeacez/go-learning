package main

import "fmt"

func main() {
	x := 0
	if x < 100 {
		fmt.Print("x 小于100")
	}

	if y := 500; y > 100 {
		fmt.Println("y 大于100")
	}

	if z := 500; z > 100 {
		fmt.Println("z 大于100")
	} else {
		fmt.Println("z 小于100")
	}
}
