package main

// go doesn't have classes
// but you can define methods on types
// a method is a function with a special receiver argument

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// this is a method
// the Abs method has a receiver of type Vertex named v
// you can only declare a method with a receiver whose type is defined in the same package as the method
// you cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int)
// what this method does is it returns the absolute value of the vertex
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
