package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// functions with pointer arguments must take a pointers
// methods with pointer receivers take either a value or a pointer as the receiver when they are called
func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)
	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)
	fmt.Println(v, p)
	// fmt.Println(v)
	// returns {6 8} with value receiver
	// returns {60 80} with pointer receiver
	// fmt.Println(p)
	// returns &{4 3} normally
	// returns &{12 9} with value receiver
	// it all works either way.
}
