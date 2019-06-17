package main

import proto "github.com/golang/protobuf/proto"
import pb "github.com/adithyvisnu/learn-protobuf/model"
import "fmt"

func main() {
	fmt.Println("Client")
	name := pb.NameResponse{
		NameRes: "Adithya Visnu",
	}

	bytes, err := proto.Marshal(&name)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(bytes)
	fmt.Println("hello")
	fmt.Println(string(bytes))

	fmt.Println([]byte("Adithya Visnu"))
}
