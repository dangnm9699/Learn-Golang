package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
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
	buf := make([]byte, 1024)
	cls()
	for {
		nbytes, rAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Println(err)
			continue
		}
		receive := string(buf[:nbytes])
		// If quit
		if strings.Split(strings.TrimSpace(receive), " ")[0] == "QUIT" {
			fmt.Printf("%s -> Disconnected\n", rAddr.String())
			continue
		}
		// Print message
		fmt.Printf("%s -> %s", rAddr.String(), receive)
		go handleConn(conn, rAddr)
	}
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
