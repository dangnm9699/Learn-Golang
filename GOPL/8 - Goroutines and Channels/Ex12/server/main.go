package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client struct {
	addr string
	ch   chan<- string
}

var (
	enter   = make(chan client)
	leave   = make(chan client)
	message = make(chan string)
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcast()
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn)
	}
}

// broadcast is a function send clients' message to all clients
func broadcast() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-message:
			for cli := range clients {
				cli.ch <- msg
			}
		case cli := <-enter:
			clients[cli] = true
			current := ""
			for cl := range clients {
				current += "\t" + cl.addr + "\n"
			}
			cli.ch <- "\nCurrent clients:\n" + current
		case cli := <-leave:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

func handleConn(c net.Conn) {
	// Init a client
	msgCh := make(chan string)
	addr := c.RemoteAddr().String()
	// Start write from client's channel to conn
	go clientWriter(c, msgCh)
	// Enter
	cl := client{addr, msgCh}
	cl.ch <- "Your lAddr is " + addr
	message <- addr + " -> connected"
	enter <- cl
	fmt.Println(addr + " -> connected")
	// Send
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		message <- c.RemoteAddr().String() + " -> " + scanner.Text()
		fmt.Println(c.RemoteAddr().String() + " -> " + scanner.Text())
	}
	// Leave
	leave <- cl
	message <- c.RemoteAddr().String() + " -> disconnected"
	fmt.Println(c.RemoteAddr().String() + " -> disconnected")
	c.Close()
}

func clientWriter(c net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(c, msg)
	}
}
