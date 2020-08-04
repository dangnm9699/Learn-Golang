package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func main() {
	// for i := range pc {
	// 	fmt.Println(pc[i])
	// }
	fmt.Println("Ex4.1")
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println(diffCount(c1, c2))
}

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func diffCount(a, b [32]uint8) int {
	var res int = 0
	for i := range a {
		res += int(pc[a[i]^b[i]])
	}
	return res
}
