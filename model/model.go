package model

import (
	"encoding/json"

	proto "github.com/golang/protobuf/proto"
)

// JSMarshal do nothing
func JSMarshal(name *NameResponse) ([]byte, error) {
	return json.Marshal(name)
}

// JSUnmarshal do nothing
func JSUnmarshal(b []byte, nr *NameResponse) error {
	return json.Unmarshal(b, nr)
}

// PBMarshal marshal structs
func PBMarshal(name *NameResponse) ([]byte, error) {
	return proto.Marshal(name)
}

// PBUnmarshal unmarshal bytes
func PBUnmarshal(b []byte, nr *NameResponse) error {
	return proto.Unmarshal(b, nr)
}
