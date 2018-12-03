package main

import "fmt"

type Age int32

func (age Age) initAge() int32 {
	if age == 0 {
		age = 10
		return int32(age)
	}
	return int32(age)
}

func main() {
	age := Age(0)
	initAge := age.initAge()
	fmt.Println(initAge)
}
