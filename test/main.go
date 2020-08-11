package main

import (
	"fmt"
	"time"
)

type res struct {
	cmd     int
	rescode int
	reason  string
}

func main() {
	a := &res{1, 400, "Bad Request"}
	start := time.Now()
	for i := 0; i < 2000; i++ {
		fmt.Println(a.cmd, a.rescode, a.reason)
	}
	fmt.Println(time.Since(start))
}
