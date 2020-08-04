package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	counts := make(map[string]int) // counts of Unicode characters
	invalid := 0                   // count of invalid UTF-8 characters
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsControl(r) {
			counts["control"]++
		}
		if unicode.IsDigit(r) {
			counts["digit"]++
		}
		if unicode.IsGraphic(r) {
			counts["graphic"]++
		}
		if unicode.IsLetter(r) {
			counts["letter"]++
		}
		if unicode.IsLower(r) {
			counts["lower"]++
		}
		if unicode.IsMark(r) {
			counts["mark"]++
		}
		if unicode.IsNumber(r) {
			counts["number"]++
		}
		if unicode.IsSpace(r) {
			counts["space"]++
		}
		if unicode.IsPrint(r) {
			counts["print"]++
		}
		if unicode.IsPunct(r) {
			counts["punct"]++
		}
		if unicode.IsSymbol(r) {
			counts["symbol"]++
		}
		if unicode.IsTitle(r) {
			counts["title"]++
		}
		if unicode.IsUpper(r) {
			counts["upper"]++
		}
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%s\t%d\n", c, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
