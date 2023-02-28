package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
		// references a variable defined outside the function body
	}
	// this function returns another function
	// the returned function closes over the variable sum to form a closure
	// this is what a closure is
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
	// what this does is:
	// 1. calls adder, which assigns 0 to sum
	// 2. returns the anonymous function
	// 3. assigns the returned function to pos
	// 4. calls adder again, which assigns 0 to sum
	// 5. returns the anonymous function
	// 6. assigns the returned function to neg
	// 7. loops through the for loop
	// 8. calls pos with i as the argument
	// 9. calls neg with -2*i as the argument
	// 10. prints the returned values from pos and neg
	// 11. loops through the for loop again
	// 12. this does it until i is no longer less than 10

	// this is also a closure
	// it references adder's sum variable, which is defined outside the function body
}
