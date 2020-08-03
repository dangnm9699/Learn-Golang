package main

import "strings"

func main() {

}

func echo1(args []string) {
	var s string
	var space string
	for _, arg := range args {
		s += space
		s += arg
		space = " "
	}
}

func echo2(args []string) {
	strings.Join(args, " ")
}
