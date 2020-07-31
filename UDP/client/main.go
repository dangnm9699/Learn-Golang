package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please import host:port and your socket")
		return
	}
	server := os.Args[1]
	socket := os.Args[2]
	skt, _ := strconv.Atoi(socket)
	fmt.Println("Connecting...")
	time.Sleep(1 * time.Second)
	rAddr, err := net.ResolveUDPAddr("udp", server)
	if err != nil {
		log.Fatal(err)
	}
	lAdrr := net.UDPAddr{Port: skt}
	conn, err := net.DialUDP("udp", &lAdrr, rAddr)
	if err != nil {
		log.Fatal(err)
	}
	cls()
	fmt.Println("Successfully connected!")
	conn.Write([]byte("Connected\n"))
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		input, _ := reader.ReadString('\n')
		conn.Write([]byte(input))
		// Print sresponse
		buf := make([]byte, 1024)
		nbytes, _, _ := conn.ReadFromUDP(buf)
		fmt.Printf("Response: %s", string(buf[:nbytes]))
		// If quit
		if strings.Split(strings.TrimSpace(input), " ")[0] == "QUIT" {
			fmt.Println("Disconnecting...")
			time.Sleep(1 * time.Second)
			cls()
			return
		}
	}
}

// Clear Screen
func cls() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
