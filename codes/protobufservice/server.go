package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type poServer struct{}

// SayHello implements helloworld.GreeterServer
func (s *poServer) Purchase(ctx context.Context, req *TransRequest) (*TransResponse, error) {
	log.Printf("SERVER: Got Order from Purchase Order Client: %v.\n\n\n", req)

	var index int
	var total int32
	statues := make(map[string]TransResponse_OrderStatue)
	for k, v := range req.Itemlist {
		statues[k] = TransResponse_ACCEPT
		total = total + req.Counts[index]*v
		index++
	}

	return &TransResponse{
		Accept:  true,
		Total:   total,
		Statues: statues}, nil
}

func doServerAction() {
	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterPurchaseOrderServiceServer(s, &poServer{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
