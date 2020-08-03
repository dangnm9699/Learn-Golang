package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	fmt.Println("Connected")
	conn.Write([]byte("Connected\n"))
	// fmt.Fprint(conn, "Connected\n")
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("send   : ")
		input, _ := reader.ReadString('\n')
		conn.Write([]byte(input + "\n"))
		// fmt.Fprintf(conn, input+"\n")
		res, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("receive: " + res)
		a := strings.Split(strings.TrimSpace(input), " ")[0]
		if a == "quit" {
			fmt.Println("Disconnected")
			return
		}
	}
}
