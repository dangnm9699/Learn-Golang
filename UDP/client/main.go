package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"

	"google.golang.org/protobuf/proto"
)

func main() {
	rAddr, err := net.ResolveUDPAddr("udp4", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialUDP("udp4", nil, rAddr)
	if err != nil {
		log.Fatal(err)
	}
	// Goroutine here
	go func() {
		for {
			sendData(conn)
		}
	}()
	for {
		readResp(conn)
	}
}
func sendData(c *net.UDPConn) {
	send := Packet{
		Cmd: rand.Int31n(2) + 1,
		Info: &User{
			Id:    os.Args[1],
			Score: rand.Int31n(10),
		},
	}
	data, err := proto.Marshal(&send)
	if err != nil {
		log.Fatal(err)
	}
	c.Write(data)
	// time.Sleep(time.Second)
}

func readResp(c *net.UDPConn) {
	buf := make([]byte, 2048)
	nbytes, _, err := c.ReadFromUDP(buf)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(buf[:nbytes]))
}
