package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/protobuf/proto"
)

var queue chan bool

func main() {
	queue = make(chan bool, 1000000)
	createFile()
	addr, err := net.ResolveUDPAddr("udp4", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		handleConn(conn)
	}
}

func handleConn(c *net.UDPConn) {
	buf := make([]byte, 2048)
	nbytes, addr, err := c.ReadFromUDP(buf)
	if err != nil {
		log.Println(err)
		return
	}
	// Goroutine here to execute
	execute(c, addr, buf[:nbytes])
}

func execute(c *net.UDPConn, addr *net.UDPAddr, data []byte) {
	queue <- true
	var receive Packet
	if err := proto.Unmarshal(data, &receive); err != nil {
		log.Println(err)
		<-queue
		return
	}
	c.WriteToUDP([]byte(addr.String()+" 200 OK"), addr)
	fmt.Println(addr.String(), receive.Cmd)
	<-queue
}
