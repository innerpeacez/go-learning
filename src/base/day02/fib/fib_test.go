package fib

import (
	"testing"
	"fmt"
)

func TestFibList(t *testing.T)  {
	a:=1
	b:=1

	fmt.Print(a)
	for index := 0; index < 10; index++ {
		fmt.Print(" ",b)
		tmp:=a
		a=b
		b= tmp + a
	}
	fmt.Println()
}

func TestExchange(t *testing.T) {
	a:=1
	b:=2
	a,b=b,a
	t.Log(a,b)
}
