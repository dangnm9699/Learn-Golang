package main

import (
	"log"
	"net"
	"os"
	"os/exec"
)

func main() {
	cls()
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

// Clear Screen
func cls() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
