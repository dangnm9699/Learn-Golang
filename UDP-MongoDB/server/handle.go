package main

import (
	"context"
	"net"

	"go.mongodb.org/mongo-driver/bson"
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
			Rescode: 409,
			Reason:  "Conflit: Duplicated record",
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
	filter := bson.M{"msisdn": bson.M{"$eq": req.Data.MSISDN}, "imsi": bson.M{"$eq": req.Data.IMSI}}
	result := User{}
	if err := collection.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		*res = Response{
			Cmd:     req.Cmd,
			Rescode: 404,
			Reason:  "Not Found",
		}
		return
	}
	udpate := bson.M{"$set": bson.M{
		"imsi": req.Data.IMSI,
		"name": stringToModify(result.Name, req.Data.Name),
		"id":   stringToModify(result.ID, req.Data.ID),
		"dob":  stringToModify(result.DOB, req.Data.DOB)}}
	_, err := collection.UpdateOne(context.TODO(), filter, udpate)
	if err != nil {
		*res = Response{
			Cmd:     req.Cmd,
			Rescode: 500,
			Reason:  "Internal Server Error",
		}
		return
	}
	*res = Response{
		Cmd:     req.Cmd,
		Rescode: 204,
		Reason:  "No Content: OK",
	}
	return
}

func delete(req *Request, res *Response) {
	filter := bson.M{"msisdn": bson.M{"$eq": req.Data.MSISDN}, "imsi": bson.M{"$eq": req.Data.IMSI}}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		*res = Response{
			Cmd:     req.Cmd,
			Rescode: 500,
			Reason:  "Internal Server Error",
		}
		return
	}
	if result.DeletedCount == 0 {
		*res = Response{
			Cmd:     req.Cmd,
			Rescode: 404,
			Reason:  "Not Found",
		}
		return
	}
	*res = Response{
		Cmd:     req.Cmd,
		Rescode: 204,
		Reason:  "No Content: OK",
	}
	return
}
