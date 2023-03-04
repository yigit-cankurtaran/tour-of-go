package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
	// created a struct with two fields, When and What
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
	// Error() method for MyError struct
	// this method is called when the error is printed using fmt.Print(e)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
		// returns a pointer to a MyError struct
	}
}

func main() {
	if err := run(); err != nil {
		// this checks for an error
		// if there is an error, it will print the error
		fmt.Println(err)
		// a nil error is success, a non-nil error is a failure.
	}
}
