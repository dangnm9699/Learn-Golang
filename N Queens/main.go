package main

import (
	"fmt"
)

var solutions int = 0
var size int
var checkCol []bool
var checkBackSlash []bool
var checkSlash []bool

/*
 * Description: main function
 */
func main() {
	fmt.Scanf("%d\n", &size)
	checkCol = make([]bool, size+1)
	checkBackSlash = make([]bool, 2*size)
	checkSlash = make([]bool, 2*size)
	try(1)
	fmt.Print(solutions)
}

func try(row int) {
	if row == size+1 {
		solutions++
		return
	}
	for col := 1; col <= size; col++ {
		if !checkCol[col] &&
			!checkSlash[row-col+size] &&
			!checkBackSlash[row+col-1] {
			// assign to be used
			checkCol[col] = true
			checkSlash[row-col+size] = true
			checkBackSlash[row+col-1] = true
			// try next row
			try(row + 1)
			// backtrack
			checkCol[col] = false
			checkSlash[row-col+size] = false
			checkBackSlash[row+col-1] = false
		}
	}
}
