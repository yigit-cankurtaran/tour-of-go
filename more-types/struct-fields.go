package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	// this is possible because X is a field of the struct Vertex
	fmt.Println(v.X)
	// prints 4
	fmt.Println(v)
	// prints {4 2}
	// it changed the value in the struct itself
}
