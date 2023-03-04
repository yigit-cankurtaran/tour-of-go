package main

import (
	"fmt"
	// "golang.org/x/tour/tree"
	// these imports don't work locally, will run on the Go playground
)

func New(k int) *tree.Tree {
	if k == 0 {
		return nil
	}
	return &tree.Tree{New(k - 1), k, New(k - 1)}
}

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	close(ch)
	Walk(t.Right, ch)
}

func Same(t1, t2 *tree.Tree) bool {
	if t1 == nil && t2 == nil {
		return true
	} else if t1 == nil || t2 == nil {
		return false
	} else if t1.Value != t2.Value {
		return false
	} else {
		return Same(t1.Left, t2.Left) && Same(t1.Right, t2.Right)
	}
}

func main() {
	ch := make(chan int)
	go Walk(New(3), ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("Same(New(1), New(1)):", Same(New(1), New(1))) // should be true
	fmt.Println("Same(New(1), New(2)):", Same(New(1), New(2))) // should be false
}
