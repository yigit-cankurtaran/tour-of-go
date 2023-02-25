package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	if x < 0 {
		return "math.Sqrt(-x) + 'i'"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(Sqrt(2), Sqrt(-4))
}
