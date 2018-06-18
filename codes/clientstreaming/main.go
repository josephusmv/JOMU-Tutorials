package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
)

func main() {
	go runServer()

	runClient()
}

func testStreamData() []*StreamData {
	data := make([]*StreamData, 10)

	rand.Seed(time.Now().UTC().UnixNano()) //used by getTestRanData()
	for i := int64(0); i < 10; i++ {
		data[i] = getTestRanData(i)
	}

	return data
}

func getTestRanData(id int64) *StreamData {
	var data StreamData
	data.ID = id
	data.Values = make([]string, 10)
	for i := 0; i < 10; i++ {
		data.Values[i] = fmt.Sprintf("This is message %d for ID: %d", i, id)
	}

	data.TimeStamp = &timestamp.Timestamp{}
	data.TimeStamp.Seconds = time.Now().Unix()
	data.TimeStamp.Nanos = int32(time.Now().Nanosecond())

	return &data
}

func randomString(len int) string {
	var result bytes.Buffer
	for i := 0; i < len; i++ {

		upperCharInt := 65 + rand.Intn(90-65)
		lowerCharInt := 97 + rand.Intn(122-97)
		punctList := []int{'@', '#', '_'}

		switch rand.Intn(30) {
		case 0:
			result.WriteString(string(punctList[0]))
		case 1:
			result.WriteString(string(punctList[1]))
		case 3:
			result.WriteString(string(punctList[2]))
		case 5, 6, 7, 8, 9, 10, 11:
			result.WriteString(string(upperCharInt))
		default:
			result.WriteString(string(lowerCharInt))
		}
	}
	return result.String()
}
