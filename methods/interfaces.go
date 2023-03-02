package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
	// an interface is a set of method signatures
	// value of interface type can hold any value that implements those methods
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f // a MyFloat implements Abser
	// myFloat is float64, so it implements Abs() float64
	a = &v // a *Vertex implements Abser
	// *Vertex is a struct, so it implements Abs() float64

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// a = v
	// gives an error, cannot use v (type Vertex) as type Abser in assignment

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
