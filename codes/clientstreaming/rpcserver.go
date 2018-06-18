package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/golang/protobuf/ptypes/wrappers"
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
func (s *serverImpl) EnvoyerStream(stream Demostream_EnvoyerStreamServer) error {
	s.currentID++
	for {
		data, err := stream.Recv()

		if err == io.EOF {
			fmt.Println("[SERVER]: \tRecieved EOF, send response.\n")
			return stream.SendAndClose(&wrappers.Int64Value{Value: s.currentID})
		}

		if err != nil {
			return err
		}

		fmt.Printf("[SERVER]: \tRecieved Stream data --- %d ---: id: %d, at: %s\n", s.currentID, data.ID, data.TimeStamp.String())
	}

	return nil
}
