package fn

import "testing"

// 可变长参数
func sum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func TestSum(t *testing.T) {
	t.Log(sum(1, 2, 3, 4, 5, 6, 7))
}
