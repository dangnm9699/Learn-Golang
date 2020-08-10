package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"runtime"
	"strconv"
	"strings"
	"time"

	"google.golang.org/protobuf/proto"
)

func main() {
	runtime.GOMAXPROCS(1)
	rAddr, err := net.ResolveUDPAddr("udp4", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialUDP("udp4", nil, rAddr)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			sendData(conn)
			time.Sleep(time.Nanosecond)
		}
	}()
	for i := 1; true; i++ {
		readResp(conn)
		fmt.Println(i)
	}
}

func randString() string {
	a := strconv.Itoa(int(rand.Int31n(1000000000)))
	a = strings.Repeat("0", 9-len(a)) + a
	return a
}

func sendData(c *net.UDPConn) {
	send := Request{
		Cmd: rand.Int31n(3) + 1,
		Data: &User{
			MSISDN: "84" + randString(),
			IMSI:   "45204",
			Name:   "dangnm",
			ID:     "125832414",
			DOB:    "090699",
		},
	}
	data, err := proto.Marshal(&send)
	if err != nil {
		log.Fatal(err)
	}
	c.Write(data)
	// time.Sleep(time.Nanosecond)
}

func readResp(c *net.UDPConn) {
	buf := make([]byte, 2048)
	nbytes, _, err := c.ReadFromUDP(buf)
	if err != nil {
		log.Println(err)
		return
	}
	res := &Response{}
	if err := proto.Unmarshal(buf[:nbytes], res); err != nil {
		log.Println(err)
		return
	}
	fmt.Println(res.Cmd, res.Rescode, res.Reason)
}
