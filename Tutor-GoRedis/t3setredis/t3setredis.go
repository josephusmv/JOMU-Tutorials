package main

import (
	"fmt"
	"reflect"

	"github.com/garyburd/redigo/redis"
)

type RedisValue struct {
	t     reflect.Type
	field string
	value interface{}
}

func NewRedisValue(field string, value interface{}) *RedisValue {
	return &RedisValue{
		field: field,
		value: value,
		t:     reflect.TypeOf(value),
	}
}

func NewRedisValueFromTypeAssertion(field string, t reflect.Type, value interface{}, err error) *RedisValue {
	rv := &RedisValue{
		field: field,
		t:     t,
	}

	switch t.Kind() {
	case reflect.Int:
		rv.value, err = redis.Int(value, err)
	case reflect.Float64:
		rv.value, err = redis.Float64(value, err)
	case reflect.String:
		rv.value, err = redis.String(value, err)
	case reflect.Bool:
		rv.value, err = redis.Bool(value, err)
	}

	return rv
}

func (rv *RedisValue) GetFieldValuePair() (string, interface{}) {
	return rv.field, rv.value
}

func main() {
	testRedisHash()
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

	rvArray := make([]*RedisValue, 5)
	for i := range fields {
		rvArray[i] = NewRedisValue(fields[i], values[i])

		reply, errSet := c.Do("HMSET", "User002", fields[i], values[i])
		if errSet != nil {
			panic(errSet)
		}

		fmt.Printf("Set Type: %v, Set Value: %v, result: %v \n",
			rvArray[i].t, rvArray[i].value, reply)
	}

	fmt.Println("\n  *********************** ")

	rsltArray := make([]*RedisValue, 5)
	for i := range fields {

		reply, errSet := c.Do("HMGET", "User002", rvArray[i].field)
		if errSet != nil {
			panic(errSet)
		}

		rsltArray[i] = NewRedisValueFromTypeAssertion(rvArray[i].field, rvArray[i].t, reply, errSet)

		fmt.Printf("Get Type: %v, Get Value: %v \n", rvArray[i].t, rvArray[i].value)
	}

}
