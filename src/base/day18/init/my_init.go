package my_init

import "fmt"

func HelloInit() {
}

// init 方法
// 一个go 文件中可以又多个 init 方法，执行顺序按导入的包顺序和init方法的顺序 从上到先 依次执行

func init() {
	fmt.Println("Hello Init")
}
