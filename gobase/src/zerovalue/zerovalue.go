package main

import "fmt"

func main() {
	var _int int
	var _int8 int8
	var _int16 int16
	var _int32 int32

	var _uint uint
	var _uint8 uint8
	var _uint16 uint16
	var _uint32 uint32
	var _uint64 uint64
	var _uintptr uintptr

	var _bool bool
	var _float32 float32
	var _float64 float64
	var _rune rune
	var _string string
	var _byte byte
	var _complex64 complex64
	var _complex128 complex128

	fmt.Println("int 类型零值为：", _int)
	fmt.Println("int8 类型零值为：", _int8)
	fmt.Println("int16 类型零值为：", _int16)
	fmt.Println("int32 类型零值为：", _int32)
	fmt.Println("-----------------------------------")
	fmt.Println("uint 类型零值为：", _uint)
	fmt.Println("uint8 类型零值为：", _uint8)
	fmt.Println("uint16 类型零值为：", _uint16)
	fmt.Println("uint32 类型零值为：", _uint32)
	fmt.Println("uint64 类型零值为：", _uint64)
	fmt.Println("uintptr 类型零值为：", _uintptr)
	fmt.Println("-----------------------------------")
	fmt.Println("bool 类型零值为：", _bool)
	fmt.Println("-----------------------------------")
	fmt.Println("float64 类型零值为：", _float64)
	fmt.Println("float32 类型零值为：", _float32)
	fmt.Println("-----------------------------------")
	fmt.Println("string 类型零值为：", _string)
	fmt.Println("-----------------------------------")
	fmt.Println("rune 类型零值为：", _rune)
	fmt.Println("-----------------------------------")
	fmt.Println("byte 类型零值为：", _byte)
	fmt.Println("-----------------------------------")
	fmt.Println("complex64 类型零值为：", _complex64)
	fmt.Println("complex128 类型零值为：", _complex128)

}
