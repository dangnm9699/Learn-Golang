package main

import (
	"io"
	"log"
	"net"
	"os"
	"os/exec"
)

func main() {
	cls()
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan bool)
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- true // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done // wait for background goroutine to finish
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

// Clear Screen
func cls() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
