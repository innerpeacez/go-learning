package _map

import "testing"

func TestMap6(t *testing.T) {
	// map 的value 可以是函数
	map1 := map[string]func(num int) int{}
	map1["num"] = func(num int) int {
		return num
	}
	map1["square"] = func(num int) int {
		return num * num
	}
	map1["cube"] = func(num int) int {
		return num * num * num
	}

	t.Logf("%d", map1["num"](2))
	t.Logf("%d", map1["square"](2))
	t.Logf("%d", map1["cube"](2))
}

func TestSet(t *testing.T) {
	// 使用 map 实现 set
	mySet := map[int]bool{}
	mySet[1] = true
	if mySet[1] {
		t.Log("key 1 is existing", mySet[1])
	} else {
		t.Log("key 1 is not existing")
	}
}
