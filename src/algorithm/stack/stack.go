package stack

import "errors"

type Stack struct {
	top int
	arr [10]int
}

func (stack *Stack) push(num int) (bool, error) {
	if len(stack.arr) > cap(stack.arr) {
		return false, errors.New("overflow")
	}
	tmp := stack.top
	stack.top++
	stack.arr[tmp] = num
	return true, nil
}

func (stack *Stack) pop() (int, error) {
	if stack.top == 0 {
		return 0, errors.New("underflow")
	}

	pop := stack.arr[stack.top]
	stack.top--
	return pop, nil
}

func (stack *Stack) empty() bool {
	if stack.top == 0 {
		return true
	}
	return false
}
