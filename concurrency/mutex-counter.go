package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter = safe to use concurrently
type SafeCounter struct {
	mu sync.Mutex
	// standard library, Lock and Unlock methods
	v map[string]int
}

// Inc = incrementer
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// **Lock so only one goroutine at a time can access the map c.v.**
	c.v[key]++
	c.mu.Unlock()
}

// Value == value of counter
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
