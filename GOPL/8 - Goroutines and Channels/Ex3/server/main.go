package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	cls()
	if len(os.Args) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	port := ":" + os.Args[1]
	listener, err := net.Listen("tcp", port)
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

// Clear Screen
func cls() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}
