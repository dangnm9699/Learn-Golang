package main

import "fmt"

func main() {
	s := []string{"dang", "nguyen", "nguyen", "minh"}
	fmt.Println(removeDuplicates(s))
}

func removeDuplicates(strings []string) []string {
	count := 0
	var s []string
	check := make(map[string]bool)
	for _, str := range strings {
		if !check[str] {
			count++
			s = append(s, str)
			check[str] = true
		}
	}
	return s[:count]
}
