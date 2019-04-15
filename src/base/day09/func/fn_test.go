package _func

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestFn(t *testing.T) {
	i, _ := returnMultiValues()

	t.Log(i)
}

// 计算函数执行的时常
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		end := inner(n)
		fmt.Println("time spent", time.Since(start).Seconds())
		return end
	}
}

func returnValue(op int) int {
	time.Sleep(time.Second * 10)
	return rand.Intn(10)
}

// 函数式编程
func TestTimeSpent(t *testing.T) {
	spent := timeSpent(returnValue)
	t.Log(spent(10))
}
