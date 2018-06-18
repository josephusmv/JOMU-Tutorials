package main

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"time"

	"github.com/golang/protobuf/ptypes/wrappers"

	"google.golang.org/grpc"
)

func runClient() {
	conn, err := grpc.Dial("localhost:5324", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//通过grpc.Dial得到的网络层连接获取一个类型为DemostreamClient的实例
	client := NewDemostreamClient(conn)

	for i := 0; i < 5; i++ {
		id := rand.Intn(10)
		if id > 9 {
			break
		}

		//根据protobuf中定义发送一个int64的ID值
		stream, err := client.RecevoirStream(ctx, &wrappers.Int64Value{Value: int64(id)})
		for { //通过EOF标志或者错误Error来退出循环
			var data *StreamData
			data, err = stream.Recv() //调用了前述的Recv()方法
			if err == io.EOF {
				break //本次的流已经接受完毕，退出循环
			}
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("[CLIENT]: \tRecieved Stream data[%d]: %s at %v]\n", data.ID, data.Values[0], data.TimeStamp)
		}
	}

}
