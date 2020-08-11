package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"runtime"
	"sync/atomic"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type el struct {
	data []byte
	addr net.Addr
}

var collection *mongo.Collection
var count int64

func main() {
	runtime.GOMAXPROCS(2)
	queue := make(chan el, 5000)
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
	// addr, err := net.ResolveUDPAddr("udp4", "localhost:8000")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	conn, err := net.ListenPacket("udp4", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			buf := make([]byte, 32*1024)
			nbytes, addr, err := conn.ReadFrom(buf)
			if err != nil {
				log.Println(err)
				continue
			}
			queue <- el{buf[:nbytes], addr}
		}
	}()
	for {
		info := <-queue
		execute(conn, info.addr, info.data)
		fmt.Printf("\r%d", atomic.AddInt64(&count, 1))
	}
}