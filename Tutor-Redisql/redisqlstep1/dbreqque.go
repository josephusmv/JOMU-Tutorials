package main

import (
	"github.com/garyburd/redigo/redis"
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

func (dbq dbReqQue)