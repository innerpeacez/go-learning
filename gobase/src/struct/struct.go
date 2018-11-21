package main

import "fmt"

type Var struct {
	X int
	Y int
}

func main() {
	fmt.Println(Var{1, 2})

	_var := Var{1, 2}
	X := _var.X
	Y := _var.Y
	fmt.Println(X)
	fmt.Println(Y)
}
