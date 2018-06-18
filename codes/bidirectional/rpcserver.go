package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

func runServer() {
	lis, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	RegisterDemostreamServer(s, &serverImpl{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type serverImpl struct {
	currentID int64
}

//DemostreamServer接口的应用层实现，接受客户端发来的流，并且返回一个int64的ID
func (s *serverImpl) EchangerStream(stream Demostream_EchangerStreamServer) error {
	s.currentID++
	for {
		data, err := stream.Recv()

		if err == io.EOF {
			fmt.Println("[SERVER]: \tRecieved EOF, send response.\n")
			return nil
		}

		if err != nil {
			return err
		}

		fmt.Printf("[SERVER]: \tRecieved Stream data: id: %d, at: %s\n", data.ID, data.TimeStamp.String())

		serverData := testStreamData()
		for _, d := range serverData {
			d.ID = data.ID
			d.Values[0] = fmt.Sprintf("Response for Client data %d", data.ID)
			fmt.Printf("[SERVER]: \nSend Stream data: id: %d, at: %s\n", d.ID, d.TimeStamp.String())
			if err := stream.Send(d); err != nil {
				return err
			}
		}
	}

	return nil
}
