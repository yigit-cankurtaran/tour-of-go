package main

import (
	"fmt"
	"strings"
)

func main() {
	//tic-tac-toe board.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
		// the types and lengths were redundant, so they were omitted
	}

	// the inputs are in the form of row and column
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
		// this is a neat way to print out the contents of a slice
	}
}
