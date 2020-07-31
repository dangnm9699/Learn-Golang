package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	fmt.Println("Commencing countdown.")
	// tick := time.Tick(1 * time.Second)
	// for countdown := 10; countdown > 0; countdown-- {
	// 	fmt.Printf("\r%-2d", countdown)
	// 	<-tick
	// }
	select {
	case <-time.After(10 * time.Second):
	case <-abort:
		fmt.Println("Abort mission!")
		return
	}
	fmt.Println("\rLaunch...")
}
