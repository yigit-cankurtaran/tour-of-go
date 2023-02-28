package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	// the function above is an anonymous function
	// anonymous functions are useful when you want to define a function inline
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	// we use the value of hypot and pass it to compute
	fmt.Println(compute(math.Pow))
}
