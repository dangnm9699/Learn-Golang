package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				log.Println(err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, fileCount := range counts {
		n := 0
		var fileString, space string
		for file, count := range fileCount {
			n += count
			fileString += space
			fileString += file
			fileString += "(" + strconv.Itoa(count) + ")"
			space = " "
		}
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, fileString)
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	file := f.Name()
	for input.Scan() {
		line := input.Text()
		if counts[line] == nil {
			counts[line] = make(map[string]int)
		}
		counts[line][file]++
	}
}
