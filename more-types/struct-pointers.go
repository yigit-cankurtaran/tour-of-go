package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	// p is a pointer to v
	p.X = 1e9
	// the normal syntax would be (*p).X = 1e9, but Go allows us to omit the * operator
	// 1e9 is 1 * 10^9
	fmt.Println(v)
	// prints {1000000000 2}
}
