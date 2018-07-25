package main

import (
	"fmt"
	"reflect"

	"github.com/garyburd/redigo/redis"
)

func main() {
	testRedisBytes()
}

func testRedisBytes() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	bytes := []byte("This is a string for bytes")
	bytes = append(bytes, 0x03, 0x45)

	reply, errSet := c.Do("LPUSH", "MYKEY", bytes)
	if errSet != nil {
		panic(errSet)
	}

	fmt.Printf("reply type: %v\n", reflect.TypeOf(reply))
	fmt.Printf("reply Value: %v \n", reply)

	reply, errSet = c.Do("RPOP", "MYKEY")
	if errSet != nil {
		panic(errSet)
	}

	fmt.Printf("reply type: %v\n", reflect.TypeOf(reply))
	fmt.Printf("reply Value: %v \n", reply)

	bytes, errS := redis.Bytes(reply, errSet)
	if errS != nil {
		panic(errSet)
	}
	fmt.Printf("Result Value: %v \n", string(bytes))

}
