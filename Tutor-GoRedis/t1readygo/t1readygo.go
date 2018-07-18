package main

import (
	"fmt"
	"reflect"

	"github.com/garyburd/redigo/redis"
)

func main() {
	testRedisHash()
}

func testRedisString() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	_, err = c.Do("SET", "mykey", "jomu")
	if err != nil {
		panic(err)
	}

	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Value type: %v\n", reflect.TypeOf(username))
	fmt.Printf("Got mykey: %v \n", username)
}

func testRedisStringExpire() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	_, err = c.Do("SETEX", "mykey", "15", "jomu")
	if err != nil {
		panic(err)
	}

	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Value type: %v\n", reflect.TypeOf(username))
	fmt.Printf("Got mykey: %v \n", username)
}

func getUserDataForTest() ([]string, []interface{}) {
	var genData []interface{}
	genData = make([]interface{}, 5)
	genData[0] = 100013
	genData[1] = "Tomy"
	genData[2] = "Male"
	genData[3] = false
	genData[4] = 10.4567
	return []string{"id", "name", "gender", "online", "score"}, genData
}

func testRedisHash() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	fields, values := getUserDataForTest()

	for i := range fields {
		field := fields[i]
		value := values[i]

		reply, errSet := c.Do("HMSET", "User002", field, value)
		if errSet != nil {
			panic(errSet)
		}

		fmt.Printf("reply type: %v\n", reflect.TypeOf(reply))
		fmt.Printf("reply Value: %v \n", reply)
	}

	for i := range fields {
		field := fields[i]

		reply, errSet := c.Do("HMGET", "User002", field)
		if errSet != nil {
			panic(errSet)
		}

		fmt.Printf("reply type: %v\n", reflect.TypeOf(reply))
		fmt.Printf("reply Value: %v \n", reply)
	}

}
