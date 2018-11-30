package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5}

	for index, value := range nums {
		fmt.Printf("index: %v ,value: %v\n", index, value)
	}
}
