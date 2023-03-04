package main

type List[T any] struct {

	// T is a type parameter
	// any means that T can be any type
	// T can be used as a type in the struct definition
	next *List[T]
	val  T
}

func New[T any]() *List[T] {
	return &List[T]{nil, zeroValT[T]()}
}

func zeroValT[T any]() T {
	// couldn't figure out how to do this without a function
	// the 2nd nil in New returned an error
	// because it couldn't infer the type of nil at compile time, we can't use zero values directly.
	// wrote this function to return the zero value of T

	var zero T
	return zero
}

func (l *List[T]) Push(v T) {
	// create a new list node, and prepend it to the list
	l.next = &List[T]{nil, v}
}

func (l *List[T]) Pop() T {

	// T can be used as a return type
	v := l.val
	l = l.next
	return v
}

func main() {

	// T can be used as a type argument
	l := New[int]()
	l.Push(10)
	l.Push(20)
	println(l.Pop())
	println(l.Pop())
}
