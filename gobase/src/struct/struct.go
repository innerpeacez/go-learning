package main

import "fmt"

type Var struct {
	X int
	Y int
}

func main() {
	fmt.Println(Var{1, 2})
}
