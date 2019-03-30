package test_type

import (
	"fmt"
	"testing"
)

type My_Int int64

func Test_type(t *testing.T) {
	var a int
	var b int64
	var my_int My_Int
	a = int(b)
	a = int(my_int)
	fmt.Println(a, b, my_int)
}
