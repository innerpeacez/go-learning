package error

import (
	"errors"
	"testing"
)

// 预置错误类型，避免使用字符串拼接
var LessThanTwoError = errors.New("n should be not less than 2")
var LargerThanHundredError = errors.New("n should be not larger than 100")

func GetFib(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LargerThanHundredError
	}
	fibs := []int{1, 1}
	for i := 2; i < n; i++ {
		fibs = append(fibs, fibs[i-2]+fibs[i-1])
	}
	return fibs, nil
}

func TestGetFib(t *testing.T) {

	t.Log(GetFib(10))

	if fibs, err := GetFib(-10); err != nil {
		t.Error(err)
	} else {
		t.Log(fibs)
	}

}
