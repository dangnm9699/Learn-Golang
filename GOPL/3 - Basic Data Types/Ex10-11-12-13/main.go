package main

import (
	"bytes"
	"fmt"
	"strings"
)

//
const (
	KB = 1024
	MB = KB * KB
	GB = MB * KB
	TB = GB * KB
	PB = TB * KB
	EB = PB * KB
	ZB = EB * KB
	YB = ZB * KB
)

func main() {
	fmt.Println(enhanceComma("12465456465456.562"))
	fmt.Println(IsAnagrams("abcd", "cab"))
}

func comma(s string) string {
	var res bytes.Buffer
	n := len(s)
	for i := 0; i < n; i++ {
		if (i%3) == (n%3) && i > 0 {
			res.WriteString(",")
		}
		res.WriteByte(s[i])
	}
	return res.String()
}

func enhanceComma(s string) string {
	var res bytes.Buffer
	dot := strings.LastIndex(s, ".")
	if dot >= 0 {
		res.WriteString(comma(s[:dot]))
		res.WriteByte(s[dot])
	}
	res.WriteString(s[dot+1:])
	return res.String()
}

// IsAnagrams reports whether two strings
// a and b are anagrams of each other
func IsAnagrams(a, b string) bool {
	map1 := make(map[byte]int)
	map2 := make(map[byte]int)
	for i := 0; i < len(a); i++ {
		map1[a[i]]++
	}
	for i := 0; i < len(b); i++ {
		map2[b[i]]++
	}
	if len(map1) != len(map2) {
		return false
	}
	for char := range map1 {
		if map1[char] != map2[char] {
			return false
		}
	}
	return true
}
