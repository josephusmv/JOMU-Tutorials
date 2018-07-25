package main

import (
	"fmt"
	"testing"
)

//go test -v -cover -coverprofile cover.out -run TestBSTInsertDelete
//go tool cover -html=cover.out -o cover.html

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

//TestDBReqQueue
//go test -v -run TestDBReqQueue
func TestDBReqQueue(t *testing.T) {
	dbq, err := newDBReqQue("tcp", "127.0.0.1:6379")
	checkError(err)

	bindVar := make([]*DBFieldValue, 2)
	bindVar[0] = &DBFieldValue{
		ValueOneof: &DBFieldValue_IntVal{IntVal: 10092},
	}
	bindVar[1] = &DBFieldValue{
		ValueOneof: &DBFieldValue_StrVal{StrVal: "n"},
	}

	dbsel := &DBSelect{
		Cols:     []string{"id", "player_name", "online", "address"},
		Where:    "id=? OR online=?",
		BindVars: bindVar,
	}

	dbact := &DBAction{
		Method: DBMethod_SELECT,
		MthdOneof: &DBAction_Select{
			Select: dbsel,
		},
	}

	key, reply, errp := dbq.pushCommand(dbact)
	checkError(errp)
	fmt.Printf("Key pushed: %d, reply: %v\n", key, reply)

	dbactUm, errUm := dbq.popCommand(key, 1)
	checkError(errUm)
	fmt.Printf("Key poped: %d, reply: %v\n", key, dbactUm)

}
