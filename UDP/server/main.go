package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"runtime"

	_ "github.com/mattn/go-sqlite3"
)

var queue chan bool
var db *sql.DB
var count int

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println("Initializing...")
	queue = make(chan bool, 50000)
	db, _ = sql.Open("sqlite3", "./database.db")
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS user (msisdn TEXT PRIMARY KEY NOT NULL, imsi TEXT NOT NULL, name TEXT NOT NULL, id TEXT NOT NULL, dob TEXT NOT NULL)")
	statement.Exec()
	fmt.Println("Done!")
	addr, err := net.ResolveUDPAddr("udp4", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		count++
		handleConn(conn)
		fmt.Printf("\r%10d", count)
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
	queue <- true
	go execute(c, addr, buf[:nbytes])
}
