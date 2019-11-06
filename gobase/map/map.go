package main

import "fmt"

func main() {
	// 定义map
	var map1 = map[string]string{}

	// 使用make方法定义
	map2 := make(map[int]string)

	// 赋值
	map1["decp"] = "this is map1"
	map2[1] = "this is map2"

	decp, ok := map1["decp"]
	if ok {
		fmt.Printf("decp:%s\n", decp)
	}

	// 使用delete方法删除map中的元素
	delete(map1, "decp")

	decp2, ok := map1["decp"]
	if ok {
		fmt.Printf("decp:%s\n", decp2)
	} else {
		fmt.Println("key不存在")
	}
}
