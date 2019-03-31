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
