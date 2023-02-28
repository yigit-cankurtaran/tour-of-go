package main

import (
	"fmt"
	"math"
)

type MyFloat float64

// you can only declare a method with a receiver whose type is defined in the same package as the method
// so we need to define a new type MyFloat that is a float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
	// returns the absolute value of the float
}

func main() {
	f := MyFloat(-math.Sqrt2)
	b := MyFloat(math.Sqrt2)
	fmt.Println(f.Abs())
	fmt.Println(b.Abs())
	// returns the same thing
}
