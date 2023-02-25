package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter a number: ")
	fmt.Scan(&num)
	fizzbuzz(num)
}

func fizzbuzz(num int) {
	for i := 1; i <= num; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
