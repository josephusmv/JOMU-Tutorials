package main

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	"google.golang.org/grpc"
)

func runClient() {
	conn, errc := grpc.Dial("localhost:12345", grpc.WithInsecure())
	if errc != nil {
		panic(errc)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := NewDemostreamClient(conn)
	stream, err := client.EchangerStream(ctx)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for {
			data, err := stream.Recv()
			if err == io.EOF {
				// read done.
				fmt.Println("[CLIENT]: \tRecieved EOF, exit client.\n")
				wg.Done()
				return
			}
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("[CLIENT]: \tRecieved Stream data: id: %d, value: %s, at: %s\n", data.ID, data.Values[0], data.TimeStamp.String())
		}
	}()

	data := testStreamData()
	for i, d := range data {
		if err := stream.Send(d); err != nil {
			fmt.Printf("[CLIENT]: \tSent data: %d\n", i)
			panic(err.Error())
		}
	}
	stream.CloseSend()
	wg.Wait()
}
