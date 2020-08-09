package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/protobuf/proto"
)

var queue chan bool

var db *sql.DB

func main() {
	db, _ = sql.Open("sqlite3", "./database.db")
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS user (id INTEGER, description TEXT)")
	statement.Exec()
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
	var receive Packet
	if err := proto.Unmarshal(data, &receive); err != nil {
		log.Println(err)
		<-queue
		return
	}
	// Switch CMD
	switch receive.Cmd {
	case 1:
	case 2:
	case 3:
	default:
	}
	//
	c.WriteToUDP([]byte(addr.String()+" 200 OK"), addr)
	fmt.Println(addr.String(), receive.Cmd)
}
