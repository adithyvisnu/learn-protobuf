package main

import (
	"encoding/json"

	pb "github.com/adithyvisnu/learn-protobuf/model"
	proto "github.com/golang/protobuf/proto"
)

// func main() {
// 	fmt.Println("Client")
// 	name := pb.NameResponse{
// 		NameRes: "Adithya Visnu",
// 	}

// 	fmt.Println(name.GetNameRes())

// 	bytes, err := proto.Marshal(&name)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println(bytes)
// 	fmt.Println("hello")
// 	fmt.Println(string(bytes))

// 	fmt.Println([]byte("Adithya Visnu"))

// 	err = proto.Unmarshal(bytes, &name)
// 	fmt.Printf("%+v", name)

// 	fmt.Println(name.GetNameRes())
// }

// JSMarshal do nothing
func JSMarshal(name *pb.NameResponse) ([]byte, error) {
	return json.Marshal(name)
}

// PBMarshal do nothing
func PBMarshal(name *pb.NameResponse) ([]byte, error) {
	return proto.Marshal(name)
}
