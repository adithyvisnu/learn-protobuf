package model

import (
	"encoding/json"

	proto "github.com/golang/protobuf/proto"
)

// JSMarshal do nothing
func JSMarshal(name *NameResponse) ([]byte, error) {
	return json.Marshal(name)
}

// PBMarshal do nothing
func PBMarshal(name *NameResponse) ([]byte, error) {
	return proto.Marshal(name)
}
