package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe bool = false
	MaxInt uint64 = 1024 >> 2
	Z complex128 = cmplx.Sqrt(25 + 10i)
)

func main (){
	fmt.Printf("Type:  %T value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T value: %v\n", Z, Z)
}