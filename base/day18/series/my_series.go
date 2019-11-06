package series

import "fmt"

func init() {
	fmt.Println("this also is init method")
}

func GetFibSerice(n int) []int {
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret
}

func init() {
	fmt.Println("this is init method")
}
