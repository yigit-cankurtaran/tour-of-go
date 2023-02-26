package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	// the above is the same as i = 21
	// the * operator denotes the pointer's underlying value
	// the & operator generates a pointer to its operand
	fmt.Println(i) // see the new value of i
	// pointers work pretty much the same as in C

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	// pointers can be printed too
	fmt.Println(&i)
	fmt.Println(&j)
	// the output is the memory address of the variable
	// the memory address is different for each run and unique for each variable
}
