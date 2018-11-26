package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// switch执行顺序，从上到下，满足条件则停止
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Println("%s.", os)
	}

	today := time.Now().Weekday()

	switch time.Saturday {
	case today:
		fmt.Println("今天")
	case today + 1:
		fmt.Println("明天")
	case today + 2:
		fmt.Println("后天")
	default:
		fmt.Println("好多天后")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
