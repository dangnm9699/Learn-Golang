package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp4", ":8000")
	checkErr(err)
	conn, err := net.ListenUDP("udp4", addr)
	checkErr(err)
	for {
		handler(conn)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func handler(c *net.UDPConn) {
	buf := make([]byte, 2048)
	nbytes, addr, err := c.ReadFromUDP(buf)
	if err != nil {
		log.Println(err)
		return
	}
	data := buf[:nbytes]
	// var receive Packet
	// if err := proto.Unmarshal(data, &receive); err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Println(addr.String(), receive.Cmd, receive.Info.Id, receive.Info.Score)
	fmt.Println(addr.String(), data)
}
