package test

import "testing"

func TestLoop(t *testing.T) {

	// go中的循环只有for一个关键字
	//for j := 0; j < 10; j++ {
	//
	//}

	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
}

func TestLoop2(t *testing.T) {
	for n := 0; n < 5; n++ {
		t.Log(n)
	}
}
