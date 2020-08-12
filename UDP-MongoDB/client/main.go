package main

import (
	"fmt"
	"log"
	"net"
	sync "sync"
	"time"

	"google.golang.org/protobuf/proto"
)

var maxsize int = 2000
var name []string = []string{
	"dangnm", "ducnt", "trungnx", "hieunm", "dongnt", "trunglq",
	"thuannt", "quanvd", "luongnv", "thanhntt", "dungbv", "sonna",
	"tuannm", "ducdm", "huyvq", "trangntt", "diunt", "khanhdt",
	"thanhndt", "haitt", "duynv", "vinhnt", "linhnt", "thudt"}

func main() {
	w := sync.WaitGroup{}
	w.Add(1)
	rAddr, err := net.ResolveUDPAddr("udp4", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialUDP("udp4", nil, rAddr)
	if err != nil {
		log.Fatal(err)
	}
	start := time.Now()
	go func() {
		for i := 0; i < maxsize; i++ {
			sendRequest(conn)
			readResponse(conn)
		}
		w.Done()
	}()
	w.Wait()
	fmt.Println(time.Since(start))
}

func sendRequest(c *net.UDPConn) {
	send := Request{
		Cmd: 1,
		Data: &User{
			MSISDN: randMSISDN(),
			IMSI:   randIMSI(),
			Name:   name[r.Int31n(24)],
			ID:     randID(),
			DOB:    randDOB(),
		},
	}
	data, err := proto.Marshal(&send)
	if err != nil {
		log.Fatal(err)
	}
	c.Write(data)
}

func readResponse(c *net.UDPConn) {
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
	// fmt.Println(res.Cmd, res.Rescode, res.Reason)
}
