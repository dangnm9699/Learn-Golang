package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Error, os.Args must have 2 arguments!")
		os.Exit(1)
	}
	counts := make(map[string]int) // counts of Unicode characters
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occurred while opening file %s", os.Args[1])
	}
	read := bufio.NewScanner(file)
	read.Split(bufio.ScanWords)
	for read.Scan() {
		counts[read.Text()]++
	}
	for k, v := range counts {
		fmt.Printf("%s\t%d\n", k, v)
	}
}
