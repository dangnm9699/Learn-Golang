package main

import (
	"fmt"
	"os"
)

var step, size, solutions int

var board [][]int

var validMoveRow = []int{2, 1, -1, -2, -2, -1, 1, 2}
var validMoveCol = []int{1, 2, 2, 1, -1, -2, -2, -1}

func main() {
	var r, c int
	fmt.Scanf("%d\n", &size)
	fmt.Scanf("%d\n", &r)
	fmt.Scanf("%d\n", &c)
	board = make([][]int, size+1)
	for i := 1; i <= size; i++ {
		board[i] = make([]int, size+1)
	}
	board[r][c] = 1
	step++
	// cls()                                           // OPTIONAL
	// go spinner()                                    // OPTIONAL
	KnightTour(r, c) // NECESSARY
	// cls()                                           // OPTIONAL
	// fmt.Printf("There are %d solutions", solutions) // OPTIONAL
}

// KnightTour is a function that check if a bishop
// an go to every places of chess board
func KnightTour(r, c int) {
	if step == size*size {
		// ALTERNATIVE
		// solutions++
		// return
		// REQUIRED
		for i := 1; i <= size; i++ {
			for j := 1; j <= size; j++ {
				fmt.Printf("%d ", board[i][j])
			}
			fmt.Print("\n")
		}
		os.Exit(0)
	}
	for i := 0; i < 8; i++ {
		nextRow := r + validMoveRow[i]
		nextCol := c + validMoveCol[i]
		if nextCol > 0 && nextCol <= size &&
			nextRow > 0 && nextRow <= size {
			if board[nextRow][nextCol] == 0 {
				step++
				board[nextRow][nextCol] = step
				KnightTour(nextRow, nextCol)
				board[nextRow][nextCol] = 0
				step--
			}
		}
	}
}

// func printSolution() {
// 	fmt.Printf("Solution: %d\n", solutions)
// 	fmt.Println(strings.Repeat("-", 3*size-1))
// 	for i := 1; i <= size; i++ {
// 		for j := 1; j <= size; j++ {
// 			fmt.Printf("%-3d", board[i][j])
// 		}
// 		fmt.Print("\n")
// 	}
// 	fmt.Println(strings.Repeat("-", 3*size-1))
// 	fmt.Print("\n\n")
// }

// func spinner() {
// 	for {
// 		for _, c := range `-\|/` {
// 			fmt.Printf("\r%c", c)
// 			time.Sleep(100 * time.Millisecond)
// 		}
// 	}
// }

// func cls() {
// 	cmd := exec.Command("cmd", "/c", "cls")
// 	cmd.Stdout = os.Stdout
// 	cmd.Run()
// }
