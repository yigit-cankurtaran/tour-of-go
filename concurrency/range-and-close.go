package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		// receives values until closed
		c <- x
		x, y = y, x+y
	}
	close(c)
	// no more values will be sent on this channel
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// cap is capacity
	for i := range c {
		fmt.Println(i)
	}
}
