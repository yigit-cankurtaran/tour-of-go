package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	// pointer of type Vertex
	// the Scale method operates on a pointer to a Vertex
	// receiver type has the literal syntax *T for some type T
	// pointer receivers are **more efficient** than value receivers
	// because they can avoid copying the value on each method call
	v.X = v.X * f
	v.Y = v.Y * f
	// what the function does is it scales the vertex by a factor of 10
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	// the scale method doesn't work properly with a value receiver
	// returns 5, instead of returning 50
	// the reason is that the Scale method must modify its receiver
	fmt.Println(Abs(v))
}
