package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool   = false
	MaxInt uint64 = 1<<64 - 1
	// uint64 is an unsigned 64-bit integer
	// max value is 2^64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
	// complex128 is a 128-bit complex number
	// real and imaginary parts are float64
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
