package main

import (
	"context"
	"net"

	"google.golang.org/protobuf/proto"
)

func execute(c *net.UDPConn, addr *net.UDPAddr, data []byte) {
	req := &Request{}
	res := &Response{}
	defer func() {
		data, _ := proto.Marshal(res)
		c.WriteToUDP(data, addr)
		// <-pool
	}()
	if err := proto.Unmarshal(data, req); err != nil {
		*res = Response{
			Cmd:     req.Cmd,
			Rescode: 400,
			Reason:  "Bad Request: Cannot unmarshal",
		}
		return
	}
	if err := validate(req); err != nil {
		*res = Response{
			Cmd:     req.Cmd,
			Rescode: 400,
			Reason:  "Bad Request: Invalid data",
		}
		return
	}
	// Switch CMD
	switch req.Cmd {
	case 1:
		// CREATE
		create(req, res)
	case 2:
		// UPDATE
		update(req, res)
	case 3:
		// DELETE
		delete(req, res)
	default:
		// METHOD NOT ALLOWED
		*res = Response{
			Cmd:     req.Cmd,
			Rescode: 405,
			Reason:  "Method Not Allowed",
		}
	}
}

func create(req *Request, res *Response) {
	_, err := collection.InsertOne(context.TODO(), req.Data)
	if err != nil {
		*res = Response{
			Cmd:     req.Cmd,
			Rescode: 400,
			Reason:  "Bad Request: Invalid data",
		}
		return
	}
	*res = Response{
		Cmd:     req.Cmd,
		Rescode: 201,
		Reason:  "Created: OK",
	}
	return
}

func update(req *Request, res *Response) {
	// _, err := collection.UpdateOne(context.TODO(), )
}

func delete(req *Request, res *Response) {

}
