package main

import (
	"fmt"

	pb "github.com/adithyvisnu/learn-protobuf/model"
)

func main() {
	fmt.Println("====================")
	name := pb.NameResponse{
		NameRes: "Adithya Visnu",
	}
	bytes, err := pb.JSMarshal(&name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("JSON")
	fmt.Println(bytes)

	fmt.Println("=====================")
	fmt.Println()

	bytes, err = pb.PBMarshal(&name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("PB")
	fmt.Println(bytes)
}
