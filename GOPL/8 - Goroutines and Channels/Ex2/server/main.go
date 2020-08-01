package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		// Wait for clients' call
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	for {
		req, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}
		if strings.Split(strings.TrimSpace(req), " ")[0] == "quit" {
			fmt.Printf("%s -> Disconnected\n", c.RemoteAddr().String())
			return
		}
		fmt.Printf("%s -> %s", c.RemoteAddr().String(), req)
		fmt.Fprint(c, "OK\n")
	}
}
