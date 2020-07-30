package main

import (
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("Please import a port!")
	}
	port := "localhost:" + os.Args[1]
	addr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go handleConn(conn)
}

func handleConn(c *net.UDPConn) {
	for {
		
	}
}
