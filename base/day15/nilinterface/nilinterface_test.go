package nilinterface

import (
	"fmt"
	"testing"
)

func DoSomething(i interface{}) {
	if i, ok := i.(int); ok {
		fmt.Println("int", i)
		return
	}
	if i, ok := i.(string); ok {
		fmt.Println("string", i)
		return
	}
	fmt.Println("Unknow Type")
}

func TestDoSomething(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
	DoSomething(true)
}

func DoSomething2(i interface{}) {

	switch v := i.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("Unknow Type")
	}
}
func TestDoSomething2(t *testing.T) {
	DoSomething2(10)
	DoSomething2("10")
	DoSomething2(true)
}
