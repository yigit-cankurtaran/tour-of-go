package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	// go keyword starts a goroutine
	// a goroutine is a lightweight thread managed by the Go runtime
	// goroutines run in the same address space, so access to shared memory must be synchronized

	say("hello")
	// if we comment out the go keyword, the program will print "hello" and then "world"
	// if we don't comment out the go keyword, the program will print "hello" and then "world" at the same time
}
