package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"
)

type server struct {
	location string
	address  string
	time     string
	timezone string
}

var servers = []*server{
	&server{"New York", "localhost:8010", "", "US/Eastern"},
	&server{"Tokyo", "localhost:8020", "", "Asia/Tokyo"},
	&server{"London", "localhost:8030", "", "Europe/London"},
}

func main() {
	for _, server := range servers {
		conn, err := net.Dial("tcp", server.address)
		if err != nil {
			log.Println(err)
			continue
		}
		defer conn.Close()
		go getData(conn, server)
	}
	// Print table
	for {
		cls()
		fmt.Printf("+%s+%s+%s+\n", strings.Repeat("-", 20), strings.Repeat("-", 20), strings.Repeat("-", 20))
		fmt.Printf("|%-20s|%-20s|%-20s|\n", "Location", "Address", "Time")
		fmt.Printf("+%s+%s+%s+\n", strings.Repeat("-", 20), strings.Repeat("-", 20), strings.Repeat("-", 20))
		for _, s := range servers {
			fmt.Printf("|%-20s|%-20s|%-20s|\n", s.location, s.address, s.time)
			fmt.Printf("+%s+%s+%s+\n", strings.Repeat("-", 20), strings.Repeat("-", 20), strings.Repeat("-", 20))
		}
		time.Sleep(1 * time.Second)
	}
}

func getData(src io.Reader, dst *server) {
	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		dst.time = scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}

func cls() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
