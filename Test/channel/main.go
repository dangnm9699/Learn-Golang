package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	go func() {
		fmt.Println("Waiting...")
		<-ch1
		fmt.Println("OK")
	}()
	time.Sleep(time.Second)
	select {
	case msg := <-ch1:
		fmt.Println("receive", msg)
	default:
		fmt.Println("receive nothing")
	}
	msg := "message"
	select {
	case ch1 <- msg:
		// Block if have no receive goroutine
		// Execute default
		fmt.Println("send", msg)
	default:
		fmt.Println("send nothing")
	}
}
