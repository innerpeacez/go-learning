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

	_var2 := Var{1, 2}

	q := _var2  //将_var2的值赋值给q
	p := &_var2 // p指向_var2

	q.X = 2

	p.Y = 3
	fmt.Println(q)
	fmt.Println(p)
	fmt.Println(_var2)
	fmt.Println(*p)

	fmt.Println(v1, px, v2, v3)
}

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // 类型为 Vertex
	v2 = Vertex{X: 1}  // Y:0 被省略
	v3 = Vertex{}      // X:0 和 Y:0
	px = &Vertex{1, 2} // 类型为 *Vertex
)
