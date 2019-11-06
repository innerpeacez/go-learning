package array

import "testing"

func TestArray(t *testing.T) {
	var arr [3]int
	arr2 := [...]int{1, 2, 3, 4, 5} // 支持 ... 代替具体length
	t.Log(arr[1], arr[2])

	t.Log(arr2[3])
}

func TestLoopArray(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}

	for idx, e := range arr { // range 类似 java foreach ,idx 索引，e 元素
		t.Log(idx, e)
	}

	for _, e := range arr { // 使用 _ 占位
		t.Log(e)
	}
}

func TestArraySub(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}
	arrSub := arr[1:3] // 数组截断
	for index, value := range arrSub {
		t.Log(index, value)
	}
}
