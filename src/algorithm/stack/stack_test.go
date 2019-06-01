package stack

import "testing"

func Test_stack(t *testing.T) {
	var arr [10]int
	stack := Stack{0, arr}
	//stack.push(10)

	for i := 0; i < 10; i++ {
		stack.push(i)
	}

	i, e := stack.pop()
	if e == nil {
		t.Log(i)
	} else {
		t.Error(e)
	}

	//i2, e2 := stack.pop()
	//if e2 == nil {
	//	t.Log(i2)
	//}else {
	//	t.Error(e2)
	//}
}
