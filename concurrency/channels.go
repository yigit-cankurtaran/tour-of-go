package main

import "fmt"

func sum(s []int, c chan int) {
	// a channel is a typed conduit through which you can send and receive values with the channel operator, <-.
	// ch <- v    // Send v to channel ch.
	// v := <-ch  // Receive from ch, and assign value to v.
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
	// the data flows in the direction of the arrow
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	// channels are typed by the values they convey
	// channels can be buffered
	// by providing the buffer length as the second argument to make, we can make a buffered channel
	// sends to a buffered channel block only when the buffer is full
	// receives block when the buffer is empty
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
	// if we remove the go keyword, the program will print "0 0 0"
	// because the two lines are running sequentially
	// if we don't comment out the go keyword, the program will print "17 -5 12"
	// the reason is that the two goroutines are running concurrently
	// when they run concurrently, the program will print "17 -5 12"
	// when they run sequentially, the program will print "0 0 0"
}
