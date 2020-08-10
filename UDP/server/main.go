package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"runtime"

	_ "github.com/mattn/go-sqlite3"
)

type el struct {
	data []byte
	addr *net.UDPAddr
}

var db *sql.DB
var count int

func main() {
	runtime.GOMAXPROCS(1)
	queue := make(chan el, 5000)
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
	go func() {
		for {
			buf := make([]byte, 32*1024)
			nbytes, addr, err := conn.ReadFromUDP(buf)
			if err != nil {
				log.Println(err)
				continue
			}
			queue <- el{buf[:nbytes], addr}
		}
	}()
	for {
		count++
		info := <-queue
		execute(conn, info.addr, info.data)
		fmt.Printf("\r%d", count)
	}
}
