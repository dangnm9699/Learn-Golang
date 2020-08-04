package main

import (
	"fmt"
	"log"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	rotate(s, 8)
	fmt.Println(s)
}

func rotate(s []int, n int) {
	var res []int
	if n > len(s) {
		log.Fatalf("Invalid n = %d\n", n)
	}
	res = make([]int, len(s), len(s))
	copy(res, s[n:])
	for i := 1; i <= n; i++ {
		res[len(s)-i] = s[n-i]
	}
	copy(s, res)
}
