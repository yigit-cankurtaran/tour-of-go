package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	// this is a type switch
	// v is the value of i
	// type is the type of i
	// unlike a normal switch, the cases specify types not values
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
		// v means value, is int
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
		// q means quoted string
	default:
		fmt.Printf("I don't know about type %T!\n", v)
		// T means type
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
	do(21.0)
	// each gives different results because of the type switch
}
