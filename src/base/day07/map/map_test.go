package _map

import "testing"

func TestMap(t *testing.T) {

	// 不同的初始化方式
	map1 := map[int]int{1: 1, 2: 1, 3: 1}
	t.Log(len(map1), map1)

	map2 := map[int]int{}
	map2[3] = 30
	t.Log(len(map2), map2)

	map3 := make(map[int]int, 10)
	t.Log(len(map3), map3)
}

func TestMap2(t *testing.T) {
	// nil 的判断方式
	map1 := map[int]int{}
	map1[1] = 10

	if v, ok := map1[2]; ok {
		t.Logf("Key 1 存在 ,value 为 %d", v)
	} else {
		t.Logf("Key 不存在 %d", v)
	}
}

func TestMap3(t *testing.T) {
	// map 的遍历
	map1 := map[int]int{1: 1, 2: 2, 3: 3, 4: 4}
	for k, v := range map1 { // k 为map的key ，v为map的value
		t.Logf("key %d, value %d", k, v)
	}
}
