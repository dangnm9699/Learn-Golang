package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type res struct {
	cmd     int
	rescode int
	reason  string
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	a := "999999999"
	fmt.Println(strings.Repeat("0", 9-len(a)))
}
