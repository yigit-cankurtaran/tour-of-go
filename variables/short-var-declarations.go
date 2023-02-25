package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	// := here is a short assignment statement
	// it can be used in place of a var declaration with implicit type
	// outside of functions, every statement begins with a keyword
	// so **we can only use := inside functions**
	fmt.Println(i, j, k, c, python, java)
}
