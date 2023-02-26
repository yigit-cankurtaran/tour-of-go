package main

import (
	"fmt"
	"os"
	"os/exec"
)

var board [3][3]string
var player string = "X"

func main() {
	initializeBoard()
	printBoard()

	for {
		// get the player's move
		fmt.Printf("Player %s, enter your move (row,column): ", player)
		var row, col int
		fmt.Scan(&row, &col)

		// check if the move is valid
		if row < 0 || row > 2 || col < 0 || col > 2 {
			fmt.Println("Invalid move. Please try again.")
			continue
		}
		if board[row][col] != "" {
			fmt.Println("That space is already occupied. Please try again.")
			continue
		}

		// make the move
		board[row][col] = player
		printBoard()

		// check if the game is over
		if checkWin() {
			fmt.Printf("Player %s wins!\n", player)
			break
		}
		if checkDraw() {
			fmt.Println("The game is a draw.")
			break
		}

		// switch to the other player
		if player == "X" {
			player = "O"
		} else {
			player = "X"
		}
	}

}

func clearScreen() {
	// this is a hack to clear the screen
	// it works on Mac and Linux, but not Windows
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func initializeBoard() {
	// initialize the board to be empty
	// i used a loop to do this, but you could also do it manually
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = ""
		}
	}
}

func printBoard() {
	// clear the screen each time we print the board
	clearScreen()
	// print the numbers on the top, making it easier for the user to enter moves
	fmt.Println("   0  1  2")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", i)
		for j := 0; j < 3; j++ {
			if board[i][j] == "" {
				fmt.Printf("   ")
			} else {
				fmt.Printf(" %s ", board[i][j])
			}
			if j < 2 {
				fmt.Printf("|")
			}
		}
		fmt.Println()
		if i < 2 {
			fmt.Println("  -----------")
		}
	}
}

func checkWin() bool {
	// check rows
	for i := 0; i < 3; i++ {
		if board[i][0] != "" && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return true
		}
	}

	// check columns
	for j := 0; j < 3; j++ {
		if board[0][j] != "" && board[0][j] == board[1][j] && board[1][j] == board[2][j] {
			return true
		}
	}

	// check diagonals
	if board[0][0] != "" && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return true
	}
	if board[0][2] != "" && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return true
	}

	return false
}

func checkDraw() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == "" {
				return false
			}
		}
	}
	return true
}
