package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"testing"
	"time"
)

//Result for 20K:Read: 0.000001590 Sec, Insert: 0.000001047 Sec

func TestMapPerf(t *testing.T) {

	runtime.GOMAXPROCS(runtime.NumCPU())
	targetMap := make(map[string]interface{})
	keylist := make([]string, 10)
	var k int

	//Prepare for map with 20K elements
	for i := 0; i < 1024*20; i++ {
		key := randomString(256)
		if randInt(0, 9)%5 == 0 && k < 9 {
			keylist[k] = key
			k++
		}
		targetMap[key] = keylist
	}

	fmt.Println("")
	for i := 0; i < 10; i++ {
		keyNew := randomString(256)
		start := time.Now().UnixNano()
		targetMap[keyNew] = targetMap
		end := time.Now().UnixNano()
		dumpTimeDiff(start, end, "Insert")
	}

	var list interface{}
	fmt.Println("")
	for i := 0; i < 10; i++ {
		start := time.Now().UnixNano()
		list = targetMap[keylist[0]]
		end := time.Now().UnixNano()
		dumpTimeDiff(start, end, "Read")
	}
	fmt.Printf("%v\n", reflect.TypeOf(list))
}

func dumpTimeDiff(start, end int64, desc string) {
	diffTime := time.Unix(0, end-start)
	diffStr := diffTime.Format(time.StampNano)
	fmt.Printf("\n%s: %d.%09d, %s\n", desc, diffTime.Second(), diffTime.Nanosecond(), diffStr)
}

func randomString(l int) string {
	var result bytes.Buffer
	var temp string
	for i := 0; i < l; {
		if string(randInt(65, 90)) != temp {
			temp = string(randInt(65, 90))
			result.WriteString(temp)
			i++
		}
	}
	return result.String()
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
