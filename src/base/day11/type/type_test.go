package _type

import (
	"fmt"
	"testing"
	"unsafe"
)

// 定义一个结构体
type Employee struct {
	Id   string
	Name string
	Age  int
}

func Point(e Employee) {
	fmt.Printf("%x\n", unsafe.Pointer(&e.Name))
}

func Point2(e *Employee) {
	fmt.Printf("%x\n", unsafe.Pointer(&e.Name))
}

func TestPoint(t *testing.T) {
	// 不使用指针，会进行内存复制
	employee := Employee{"0", "zhangsan", 10}
	Point(employee)
	Point(employee)

	// 使用指针不会进行内存复制
	Point2(&employee)
	Point2(&employee)
}

func TestStruct(t *testing.T) {
	// 结构体的初始化
	employee := Employee{"0", "zhangsan", 10}
	employee2 := Employee{Name: "lisi", Age: 12}
	employee3 := new(Employee)
	employee3.Id = "3"
	employee3.Name = "wangwu"
	employee3.Age = 18

	t.Log(employee)
	t.Log(employee2)
	t.Log(employee3)
	t.Logf("%T", *employee3)
}
