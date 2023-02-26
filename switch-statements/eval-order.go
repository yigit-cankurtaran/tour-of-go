package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	// time.Now().Weekday() returns a time.Weekday, which is an int
	// 0 = Sunday, 1 = Monday, 2 = Tuesday, 3 = Wednesday, 4 = Thursday, 5 = Friday, 6 = Saturday
	// so, today is a number we can use to compare to the position of saturday
	switch time.Saturday {
	// evaluation order is from top to bottom, stopping when a case succeeds
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
		// today is sunday, so this is what i get
	}
}
