package main

import (
	"fmt"
	"net"

	"google.golang.org/protobuf/proto"
)

func execute(c *net.UDPConn, addr *net.UDPAddr, data []byte) {
	req := &Request{}
	res := &Response{}
	defer func() {
		data, _ := proto.Marshal(res)
		c.WriteToUDP(data, addr)
	}()
	if err := proto.Unmarshal(data, req); err != nil {
		res.Rescode = 400
		res.Reason = "Bad Request: " + fmt.Sprintf("%v", err)
		return
	}
	// Switch CMD
	switch req.Cmd {
	case 1:
		// CREATE
		create(req, res)
	case 2:
		// UPDATE
	case 3:
		// DELETE
	default:
		// METHOD NOT ALLOWED
	}
}

func create(req *Request, res *Response) {
	stm, _ := db.Prepare("INSERT INTO user (msisdn, imsi, name, id, dob) VALUES (?, ?, ?, ?, ?)")
	_, err := stm.Exec(req.Data.MSISDN, req.Data.IMSI, req.Data.Name, req.Data.ID, req.Data.DOB)
	if err != nil {
		*res = Response{
			Cmd:     req.Cmd,
			Rescode: 409,
			Reason:  "Conflict: Duplicated record",
		}
	}
	*res = Response{
		Cmd:     req.Cmd,
		Rescode: 201,
		Reason:  "Created: OK",
	}
}

func update(req *Request, res *Response) {
	stm, _ := db.Prepare("UDPATE user SET ")
}

func delete(req *Request, res *Response) {

}
