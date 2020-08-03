package main

import (
	"fmt"
)

var pc [256]byte

func main() {
	fmt.Printf("PopCount  = %v\n", popCount(123))
	fmt.Printf("CountEx23 = %v\n", countEx23(123))
	fmt.Printf("CountEx24 = %v\n", countEx24(123))
	fmt.Printf("CountEx25 = %v\n", countEx25(123))
}

// 10 = 1010
// 9  = 1001
// x = 0

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func popCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func countEx23(x uint64) int {
	var res int = 0
	for i := 0; i < 8; i++ {
		res += int(pc[byte(x>>(i*8))])
	}
	return res
}

func countEx24(x uint64) int {
	var res int = 0
	for i := 0; i < 64; i++ {
		if x&1 == 1 {
			res++
		}
		x = x >> 1
	}
	return res
}

func countEx25(x uint64) int {
	var res int = 0
	for x != 0 {
		x = x & (x - 1)
		res++
	}
	return res
}
