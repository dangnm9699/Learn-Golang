package main

import (
	"fmt"
	"time"
)

func main() {
	go sayHi()
	go sayHello()
	time.Sleep(time.Second)
}

func sayHi() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hi!")
		time.Sleep(time.Millisecond)
	}
}
func sayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello!")
		time.Sleep(time.Millisecond)
	}
}
