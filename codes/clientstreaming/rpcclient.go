package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
)

func runClient() {
	conn, err := grpc.Dial("localhost:12345", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := NewDemostreamClient(conn)
	data := testStreamData()

	for k := 3; k > 0; k-- { //调用 DemostreamClient.EnvoyerStream
		//为了演示，这里使用了context.TODO，
		//函数返回的是一个Demostream_EnvoyerStreamClient的流接口
		stream, errr := client.EnvoyerStream(context.TODO())
		if errr != nil {
			log.Fatalf("error call client.Purchase: %v", errr)
		}

		for i := int64(0); i < 10; i++ {
			fmt.Printf("[CLIENT]: \tSend Stream data: %d\n", data[i].ID)
			stream.Send(data[i])
			time.Sleep(time.Millisecond * 100)
		}

		time.Sleep(time.Millisecond * 500)
		reply, err2 := stream.CloseAndRecv()
		fmt.Printf("[CLIENT]: \tRecieved Stream data: %v, %v\n\n\n\n", reply, err2)
	}
}
