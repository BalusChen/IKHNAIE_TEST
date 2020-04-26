package main

import (
	"context"
	"net"

	. "github.com/BalusChen/IKHNAIE_TEST/grpc"
	"google.golang.org/grpc"
)

type HelloServiceImpl struct{}

func (*HelloServiceImpl) Hello(ctx context.Context, req *GetHelloRequest) (*GetHelloResponse, error) {
	resp := new(GetHelloResponse)
	resp.Reply = "Hello " + req.Greeting

	return resp, nil
}

func main() {
	grpcServer := grpc.NewServer()
	RegisterGetHelloServiceServer(grpcServer, new(HelloServiceImpl))

	listener, err := net.Listen("tcp", ":9877")
	if err != nil {
		panic(err)
	}
	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
