package main

import "fmt"

type People struct {
	Name string
	Age  int32
}

func (people *People) printPeople() {
	fmt.Printf("name: %s\n", people.Name)
	fmt.Printf("age: %v\n", people.Age)
}

func (people *People) updatePeople() {
	people.Age = 18
	people.Name = "zhw"
}

func main() {
	people := People{"innerpeacez", 14}
	people.printPeople()
	people.updatePeople()
	people.printPeople()
}
