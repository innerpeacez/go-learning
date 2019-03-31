package test_type

import (
	"fmt"
	"math"
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

func Test_MAX(t *testing.T) {
	var a = math.MaxInt64
	fmt.Println("a=", a)

}

func TestPoint(t *testing.T) {
	// 不支持指针运算
	a := 1
	aPtr := &a

	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	// 字符串默认值是"" 空字符串 而不是 nil
	var s string

	t.Log("*" + s + "*")
	t.Log(len(s))
}
