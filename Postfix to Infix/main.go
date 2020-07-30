package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

type stack []string

func (s *stack) push(str string) {
	*s = append(*s, str)
}

func (s *stack) pop() (elm string, ok bool) {
	if len(*s) == 0 {
		return
	}
	elm = (*s)[len(*s)-1]
	ok = true
	*s = (*s)[:len(*s)-1]
	return
}

func byteToString(b byte) string {
	tmp := ""
	buf := bytes.NewBufferString(tmp)
	buf.WriteByte(b)
	return buf.String()
}

func toInfix(postfix string) (string, error) {
	var s stack
	l := len(postfix)
	for i := 0; i < l; i++ {
		if (postfix[i] <= 'z' && postfix[i] >= 'a') ||
			(postfix[i] <= 'Z' && postfix[i] >= 'A') {
			s.push(byteToString(postfix[i]))
		} else {
			se, ok := s.pop()
			if !ok {
				return "", fmt.Errorf("Invalid postfix expression")
			}
			fi, ok := s.pop()
			if !ok {
				return "", fmt.Errorf("Invalid postfix expression")
			}
			s.push("(" + fi + byteToString(postfix[i]) + se + ")")
		}
	}
	infix, _ := s.pop()
	return infix, nil
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Please enter one and only postfix expression!")
	}
	postfix := os.Args[1]
	infix, err := toInfix(postfix)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(infix)
}
