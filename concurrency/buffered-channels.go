package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	// a buffered channel accepts a limited number of values without a corresponding receiver for those values
	ch <- 1
	ch <- 2
	// ch <- 3
	// fatal error: all goroutines are asleep - deadlock!
	// this is because the channel is full
	// the channel size is declared when the channel is created
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
