package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"time"
)

type server struct {
	location string
	address  string
	time     string
	timezone string
}

var servers = []*server{
	&server{"New York", "localhost:8010", "", "US/Eastern"},
	&server{"Tokyo", "localhost:8020", "", "Asia/Tokyo"},
	&server{"London", "localhost:8030", "", "Europe/London"},
}

func main() {
	for _, server := range servers {
		go broadcast(server)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

	}
}

func broadcast(s *server) {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Println(err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn, s)
	}
}

func handleConn(c net.Conn, s *server) {
	defer c.Close()
	for {
		location, _ := time.LoadLocation(s.timezone)
		_, err := io.WriteString(c, time.Now().In(location).Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
