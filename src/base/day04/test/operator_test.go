package test

import "testing"

func TestCompareArray(t *testing.T) {
	// 数组的比较
	// go中支持同纬度的数组进行比较

	a := [...]int{1, 2, 3, 4}
	//b:=[...]int{1,2,3,4,5}
	c := [...]int{1, 2, 3, 4}
	d := [...]int{1, 2, 4, 2}

	//t.Log(a==b)
	t.Log(a == c)
	t.Log(a == d)
}

func TestBitClear(t *testing.T) {
	// 按位清零，将符号左边数的二进制位，相同位清零，不同位保留
	a := 7 // 0111

	t.Log(a &^ 1) // 0110

	t.Log(a &^ 13)
}
