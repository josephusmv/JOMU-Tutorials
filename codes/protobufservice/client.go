package main

import (
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func makeTransactionRequest() *TransRequest {
	itemsMap := make(map[string]int32)
	itemsMap["wheel"] = 220
	itemsMap["engine"] = 3000
	itemsMap["frontlight"] = 30
	itemsMap["bearing"] = 1000
	trsReq := TransRequest{
		Buyername: "DasAuto",
		Counts:    []int32{1, 3, 4, 2},
		Itemlist:  itemsMap}

	return &trsReq
}

func doClientAction() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	poClt := NewPurchaseOrderServiceClient(conn)
	trsReq := makeTransactionRequest()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := poClt.Purchase(ctx, trsReq)
	if err != nil {
		log.Fatalf("could not get response, error: %v", err)
	}
	log.Printf("CLIENT: Got response from Purchase Order Server: %v.\n\n\n", r)
}
