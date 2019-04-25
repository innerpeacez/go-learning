package _interface

import (
	"fmt"
	"testing"
)

type Code string
type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {
}

type JavaProgrammer struct {
}

func (Go *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World\")"
}

func (Java *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.println(\"Hello World\")"
}

func TestProgrammer(t *testing.T) {
	goProgrammer := new(GoProgrammer)
	javaProgrammer := new(JavaProgrammer)

	WriteFirstProgrammer(goProgrammer)
	WriteFirstProgrammer(javaProgrammer)
}

func WriteFirstProgrammer(p Programmer) {
	fmt.Printf("%T %v \n", p, p.WriteHelloWorld())
}
