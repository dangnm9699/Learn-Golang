package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"time"
)

var port = flag.String("port", "8000", "Port number")
var tz = flag.String("tz", "US/Eastern", "Timezone")

func main() {
	flag.Parse()
	clearScreen()
	listener, err := net.Listen("tcp", "localhost:"+*port)
	fmt.Printf("TCP on %s\n", "localhost:"+*port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		location, _ := time.LoadLocation(*tz)
		_, err := io.WriteString(c, time.Now().In(location).Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
