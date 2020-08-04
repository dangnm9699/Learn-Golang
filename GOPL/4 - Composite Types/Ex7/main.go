package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Nguyễn minh Đăng"
	fmt.Println(string(reverseUTF8([]byte(s))))
}

func reverseUTF8(s []byte) []byte {
	for i := 0; i < len(s); {
		_, size := utf8.DecodeRune(s[i:])
		reverse(s[i : i+size])
		i += size
	}
	reverse(s)
	return s
}

func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
