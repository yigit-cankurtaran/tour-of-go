package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2} // has type Vertex
	v2 = Vertex{X: 1} // Y:0 is implicit
	// default int value is 0
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
	// p is a pointer to a Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
	// prints {1 2} &{1 2} {1 0} {0 0}
}
