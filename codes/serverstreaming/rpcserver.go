package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/golang/protobuf/ptypes/wrappers"
)

func runServer() {
	lis, err := net.Listen("tcp", ":5324")
	if err != nil {
		panic(err.Error())
	}

	s := grpc.NewServer()

	RegisterDemostreamServer(s, &serverImpl{})

	if err := s.Serve(lis); err != nil {
		panic(err.Error())
	}
}

type serverImpl struct {
}

//实现自动生成的接口：DemostreamServer，注意id和stream都是从流中填充。
//因为流通过这样的方式透明了相关的数据反序列化操作，在业务应用层用户得到的是具体的结构数据，
//实现完全从序列化中解脱出来
func (s *serverImpl) RecevoirStream(id *wrappers.Int64Value, stream Demostream_RecevoirStreamServer) error {
	//一个用于生成测试数据slice的测试函数
	dataArray := testStreamData()
	fmt.Printf("[SERVER]: \tRecieved for data from: %d\n", id)

	//针对本次请求，发送十个数据给客户端，之后本次流式RPC调用才完成
	for i := id.Value; i < 10; i++ {
		//调用demostreamRecevoirStreamServer.Send。
		if err := stream.Send(dataArray[i]); err != nil {
			return err
		}
	}

	return nil
}
