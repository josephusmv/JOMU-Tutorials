package main

import (
	"fmt"
	"testing"

	proto "github.com/golang/protobuf/proto"
)

type PrototBufType interface {
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func getProtoMessage() proto.Message {
	pt := &PlayersTable{}
	pt.Id = 123456
	pt.PlayerName = "Kitty Nick"
	pt.Online = "Y"
	pt.Address = "Test addr 123456"

	return pt
}

func TestProtoAPI(t *testing.T) {
	pt := getProtoMessage()

	//The data could be used in any place does not need to know the structure.
	data, err := proto.Marshal(pt)
	checkError(err)

	fmt.Printf("%v", string(data))
	/*
		var newPT proto.Message
		err = proto.Unmarshal(data, newPT)*/
}
