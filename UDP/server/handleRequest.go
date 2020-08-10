package main

import (
	"net"

	"google.golang.org/protobuf/proto"
)

func execute(c *net.UDPConn, addr *net.UDPAddr, data []byte) {
	req := &Request{}
	res := &Response{}
	defer func() {
		data, _ := proto.Marshal(res)
		c.WriteToUDP(data, addr)
		<-queue
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
	}
}

func create(req *Request, res *Response) {
	stm, _ := db.Prepare("INSERT INTO user (msisdn, imsi, name, id, dob) VALUES (?, ?, ?, ?, ?)")
	_, err := stm.Exec(req.Data.MSISDN, req.Data.IMSI, req.Data.Name, req.Data.ID, req.Data.DOB)
	if err != nil {
		*res = Response{
			Cmd:     req.Cmd,
			Rescode: 409,
			Reason:  "Conflict",
		}
	}
	*res = Response{
		Cmd:     req.Cmd,
		Rescode: 201,
		Reason:  "Created: OK",
	}
}

func update(req *Request, res *Response) {
	tx, err := db.Begin()
	if err != nil {
		*res = Response{
			Cmd:     req.Cmd,
			Rescode: 400,
			Reason:  "Bad Request",
		}
		return
	}
	rows, err := tx.Query("SELECT * FROM user WHERE msisdn = %s AND imsi = %s", req.Data.MSISDN, req.Data.IMSI)
	if err != nil {
		*res = Response{
			Cmd:     req.Cmd,
			Rescode: 404,
			Reason:  "Not Found",
		}
		tx.Rollback()
		return
	}
	var (
		MSISDN string
		IMSI   string
		Name   string
		ID     string
		DOB    string
	)
	for rows.Next() {
		rows.Scan(&MSISDN, &IMSI, &Name, &ID, &DOB)
	}
	stm, _ := tx.Prepare("UPDATE user SET msisdn = ?, imsi = ?, name = ?, id = ?, dob = ? WHERE msisdn = ?")
	_, err = tx.Stmt(stm).Exec(MSISDN, IMSI, stringToModify(req.Data.Name, Name), stringToModify(req.Data.ID, ID), stringToModify(req.Data.DOB, DOB))
	if err != nil {
		*res = Response{
			Cmd:     req.Cmd,
			Rescode: 500,
			Reason:  "Internal Server Error",
		}
		tx.Rollback()
		return
	}
	*res = Response{
		Cmd:     req.Cmd,
		Rescode: 204,
		Reason:  "No Content: OK",
	}
	tx.Commit()
}

func delete(req *Request, res *Response) {
	stm, _ := db.Prepare("DELETE FROM user WHERE msisdn = ? AND imsi = ?")
	result, err := stm.Exec(req.Data.MSISDN, req.Data.IMSI)
	if err != nil {
		*res = Response{
			Cmd:     req.Cmd,
			Rescode: 500,
			Reason:  "Intertal Server Error",
		}
		return
	}
	if nrows, _ := result.RowsAffected(); nrows == 0 {
		*res = Response{
			Cmd:     req.Cmd,
			Rescode: 400,
			Reason:  "Not Found",
		}
		return
	}
	*res = Response{
		Cmd:     req.Cmd,
		Rescode: 204,
		Reason:  "No Content: OK",
	}
}
