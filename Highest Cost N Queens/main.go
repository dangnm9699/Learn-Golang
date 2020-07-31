package main

import (
	"fmt"
)

var solutions int = 0
var max int = 0
var board [][]int
var size int
var checkCol []bool
var checkBackSlash []bool
var checkSlash []bool

func main() {
	fmt.Scanf("%d\n", &size)
	board = make([][]int, size+1)
	for i := 1; i <= size; i++ {
		board[i] = make([]int, size+1)
		for j := 1; j <= size; j++ {
			fmt.Scanf("%d", &board[i][j])
			if board[i][j] > max {
				max = board[i][j]
			}
		}
		fmt.Scanf("\n")
	}
	for i := 1; i <= size; i++ {
		for j := 1; j <= size; j++ {
			fmt.Printf("%-3d", board[i][j])
		}
		fmt.Print("\n")
	}
	checkCol = make([]bool, size+1)
	checkBackSlash = make([]bool, 2*size)
	checkSlash = make([]bool, 2*size)
	try(1, 0)
	fmt.Print(solutions)
}

func try(row, val int) {
	if row == size+1 {
		if val > solutions {
			solutions = val
		}
		return
	}
	for col := 1; col <= size; col++ {
		if !checkCol[col] &&
			!checkSlash[row-col+size] &&
			!checkBackSlash[row+col-1] &&
			val+board[row][col]+(size-row)*max > solutions {
			// assign to be used
			checkCol[col] = true
			checkSlash[row-col+size] = true
			checkBackSlash[row+col-1] = true
			// try next row
			try(row+1, val+board[row][col])
			// backtrack
			checkCol[col] = false
			checkSlash[row-col+size] = false
			checkBackSlash[row+col-1] = false
		}
	}
}
