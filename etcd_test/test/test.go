package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc_test/proto"
	"net"
)

type checkService struct {
}

func (sf *checkService) SearchCheck(ctx context.Context, request *pb.CheckRequest) (*pb.CheckResponse, error) {
	return &pb.CheckResponse{Msg: "Response test"}, nil
}

func main() {

	grpcServer := grpc.NewServer()
	pb.RegisterCheckServiceServer(grpcServer, &checkService{})

	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		panic("服务器启动失败")
	}
	grpcServer.Serve(lis)
}
