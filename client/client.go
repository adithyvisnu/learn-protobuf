package main

import (
	"fmt"
	"log"
	"net/rpc"

	pb "github.com/adithyvisnu/learn-protobuf/model"
)

const serverAddress = "0.0.0.0"

// UserParams created to be used on next development
type UserParams struct {
	UserID string
}

// Response used for global response protobuf
type Response struct {
	Res []byte
}

func main() {
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Synchronous call
	args := &UserParams{
		UserID: "2",
	}
	// Asynchronous call
	u := new(Response)
	getUserDataCall := client.Go("User.GetUserData", args, u, nil)
	replyCall := <-getUserDataCall.Done // will be equal to divCall
	// check errors, print, etc.
	log.Print(replyCall)
	if replyCall.Error != nil {
		log.Fatal(replyCall.Error)
	}
	b := replyCall.Reply.(*Response)

	userData := &pb.NameResponse{}
	fmt.Printf("%+v", b)
	err = pb.PBUnmarshal(b.Res, userData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(userData)
}
