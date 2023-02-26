package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	// no condition means true, default case
	// can be used to replace if/else
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
		//	1pm, so this is what i get
	default:
		fmt.Println("Good evening.")
	}
}
