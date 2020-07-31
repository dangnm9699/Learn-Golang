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
}

func main() {
	servers, err := getServers(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
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
		clearScreen()
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

func getServers(terms []string) ([]*server, error) {
	servers := []*server{}
	for _, term := range terms {
		info := strings.Split(term, "=")
		if len(info) != 2 {
			return nil, fmt.Errorf("Invalid format")
		}
		servers = append(servers, &server{info[0], info[1], ""})
	}
	return servers, nil
}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
