package main

import (
	"bytes"
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := "   nguyen    minh    dang   "
	fmt.Println(string(removeDuplicateSpace([]byte(s))))
}

func removeDuplicateSpace(s []byte) []byte {
	var res bytes.Buffer
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(string(s))
		if unicode.IsSpace(r) {
			r, _ := utf8.DecodeRuneInString(string(s[size:]))
			if !unicode.IsSpace(r) {
				res.WriteByte(' ')
			}
		} else {
			res.WriteRune(r)
		}
		s = s[size:]
	}
	return res.Bytes()
}
