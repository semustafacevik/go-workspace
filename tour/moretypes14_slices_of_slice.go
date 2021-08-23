package tour

import (
	"fmt"
	"strings"
)

// https://tour.golang.org/moretypes/14
// Slices of slices

// Slices can contain any type,
// including other slices.
func SlicesOfSlice() {
	// Create a tic-tac-toe board.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[1][1] = "X"
	board[2][2] = "O"
	board[0][2] = "X"
	board[2][0] = "O"
	board[2][1] = "X"
	board[1][0] = "O"
	board[0][1] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
