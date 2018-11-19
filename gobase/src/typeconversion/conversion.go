package main

import "fmt"

func main() {

	var _int int = 10
	var _float float64 = float64(_int)
	var _uint uint = uint(_float)

	_int2 := 10
	_float2 := float64(_int2)
	_uint2 := uint(_float2)

	_int3 := 10
	_float3 := _int3
	_uint3 := _float3

	fmt.Println(_int, _float, _uint, _int2, _float2, _uint2, _int3, _float3, _uint3)
}
