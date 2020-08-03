package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // connection aborted
			continue
		}
		go handleConn(conn)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	ch := make(chan bool)
	input := bufio.NewScanner(c)
	go func() {
		for {
			// If scan return true
			if input.Scan() {
				ch <- true
			} else {
				// If scan return false, client stop Stdin
				close(ch)
				return
			}
		}
	}()
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				c.Close()
				return
			}
			go echo(c, input.Text(), 500*time.Microsecond)
		case <-time.After(10 * time.Second):
			c.Close()
		}
	}
}
