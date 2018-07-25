package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
	proto "github.com/golang/protobuf/proto"
)

type dbReqQue struct {
	conn redis.Conn
}

//call redis.Dial("tcp", "127.0.0.1:6379"): proto "tcp", addr: "127.0.0.1:6379"
func newDBReqQue(proto string, addr string) (*dbReqQue, error) {
	c, err := redis.Dial(proto, addr)
	if err != nil {
		return nil, err
	}

	return &dbReqQue{conn: c}, nil
}

func (dbq *dbReqQue) genCommandKey() int64 {
	return int64(1234567890)
}

func (dbq *dbReqQue) pushCommand(dbact *DBAction) (int64, interface{}, error) {
	data, err := proto.Marshal(dbact)
	if err != nil {
		return -1, nil, err
	}

	key := dbq.genCommandKey()

	reply, errdo := dbq.conn.Do("LPUSH", key, data)

	return key, reply, errdo
}

func (dbq *dbReqQue) popCommand(key int64, timeout int) (*DBAction, error) {

	reply, errdo := dbq.conn.Do("RPOP", key) //BRPOP
	if errdo != nil {
		return nil, errdo
	}

	/*
		valStr, errtype := redis.String(reply, errdo)
		if errtype != nil {
			return nil, fmt.Errorf("Convert Command to Byte failed, error: %v,  orginal reply: %v ", errtype, reply)
		}

		var bytes []byte
		bytes = []byte(valStr)
	*/
	bytes, errtype := redis.Bytes(reply, errdo)
	if errtype != nil {
		return nil, fmt.Errorf("Convert Command to Byte failed, error: %v,  orginal reply: %v ", errtype, reply)
	}

	var dbact DBAction
	errm := proto.Unmarshal(bytes, &dbact)
	if errm != nil {
		return nil, errm
	}

	return &dbact, nil
}
