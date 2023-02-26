package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
	// prints {1 2}

	// fmt.Println(Vertex{"a", "b"})
	// returns error
	// cannot use "a" (untyped string constant) as int value in struct literal

}
