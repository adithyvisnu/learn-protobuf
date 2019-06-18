package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	pb "github.com/adithyvisnu/learn-protobuf/model"
)

// UserParams created to be used on next development
type UserParams struct {
	UserID string
}

// Response used for global response protobuf
type Response struct {
	Res []byte
}

// User created to be used on next development
type User struct{}

func main() {
	u := new(User)
	rpc.Register(u)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	log.Print("Serve at tcp :1234")
	http.Serve(l, nil)
}

// GetUserData returns user data with protobuf compression
func (*User) GetUserData(p *UserParams, res *Response) error {
	name := pb.NameResponse{}
	if p.UserID == "1" {
		name.NameRes = "Adithya Visnu"
	} else {
		name.NameRes = "Prasetyo Putra"
	}
	bytes, err := pb.PBMarshal(&name)
	if err != nil {
		fmt.Println(err)
	}
	res.Res = bytes
	return nil
}

// GetUserDataJSON return user data with json compression
func (*User) GetUserDataJSON(p *UserParams, res *Response) error {
	name := pb.NameResponse{}
	if p.UserID == "1" {
		name.NameRes = "Adithya Visnu"
	} else {
		name.NameRes = "Prasetyo Putra"
	}
	bytes, err := pb.JSMarshal(&name)
	if err != nil {
		fmt.Println(err)
	}
	res.Res = bytes
	return nil
}
