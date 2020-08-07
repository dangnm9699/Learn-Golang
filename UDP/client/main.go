package main

import (
	"log"
	"net"
)

func main() {
	rAddr, err := net.ResolveUDPAddr("udp4", ":8000")
	checkErr(err)
	conn, err := net.DialUDP("udp4", nil, rAddr)
	checkErr(err)
	for {
		sendData(conn)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func sendData(c *net.UDPConn) {
	// send := Packet{
	// 	Cmd: rand.Int31n(2) + 1,
	// 	Info: &User{
	// 		Id:    os.Args[1],
	// 		Score: rand.Int31n(10),
	// 	},
	// }
	// data, err := proto.Marshal(&send)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	c.Write([]byte("HELLO"))
	//time.Sleep(time.Second)
}
