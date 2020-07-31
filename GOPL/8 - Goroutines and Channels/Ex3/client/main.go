package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	c, ok := conn.(*net.TCPConn)
	if !ok {
		log.Fatal("Not TCP Connection")
	}
	done := make(chan bool)
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- true // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	c.CloseWrite()
	// conn.Close()
	<-done // wait for background goroutine to finish
	// conn.CloseRead()
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
