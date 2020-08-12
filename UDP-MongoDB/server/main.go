package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync/atomic"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type el struct {
	data []byte
	addr *net.UDPAddr
}

var collection *mongo.Collection
var count int64

func main() {
	// Connect DB
	clientsOpt := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientsOpt)
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	// Connect collection
	collection = client.Database("test").Collection("user")
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Disconnect to MongoDB")
	}()
	//
	addr, err := net.ResolveUDPAddr("udp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.ListenUDP("udp", addr)
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
