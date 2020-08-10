package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/mattn/go-sqlite3"
)

var queue chan bool
var db *sql.DB

func init() {
	db, _ = sql.Open("sqlite3", "./database.db")
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS user (msisdn PRIMARY KEY TEXT, imsi TEXT, name TEXT, id TEXT, dob TEXT)")
	statement.Exec()
}

func main() {
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
	buf := make([]byte, 32*1024)
	nbytes, addr, err := c.ReadFromUDP(buf)
	if err != nil {
		log.Println(err)
		return
	}
	// Goroutine here to execute
	execute(c, addr, buf[:nbytes])
}
