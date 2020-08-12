package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"sync/atomic"

	_ "github.com/mattn/go-sqlite3"
)

type el struct {
	data []byte
	addr *net.UDPAddr
}

var db *sql.DB
var count int64

func main() {
	db, _ = sql.Open("sqlite3", "./database.db")
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS user (msisdn TEXT PRIMARY KEY NOT NULL, imsi TEXT NOT NULL, name TEXT NOT NULL, id TEXT NOT NULL, dob TEXT NOT NULL)")
	statement.Exec()
	db.Exec("PRAGMA synchronous = OFF")
	fmt.Println("Initialized")
	addr, err := net.ResolveUDPAddr("udp4", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		buf := make([]byte, 32*1024)
		nbytes, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Println(err)
			continue
		}
		go execute(conn, addr, buf[:nbytes])
		fmt.Printf("\r%d", atomic.AddInt64(&count, 1))
	}
}
