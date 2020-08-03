package main

import (
	"fmt"
	"strings"
)

var args = []string{"nguyen", "minh", "dang"}

func main() {
	fmt.Println(echo1(args))
	fmt.Println(echo2(args))
}

func echo1(args []string) string {
	var s string
	var space string
	for _, arg := range args {
		s += space
		s += arg
		space = " "
	}
	return s
}

func echo2(args []string) string {
	s := strings.Join(args, " ")
	return s
}
