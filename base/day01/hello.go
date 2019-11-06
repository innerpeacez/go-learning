package main // 包名

import ( // 导包
	"os"
	"fmt"
)

func main() { // main 方法
	fmt.Println("hello go");

	if len(os.Args)>1 {
		fmt.Println("arg 1",os.Args)
	}

	os.Exit(0)
}