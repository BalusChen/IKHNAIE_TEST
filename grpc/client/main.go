package main

import (
	"context"
	"fmt"

	. "github.com/BalusChen/IKHNAIE_TEST/grpc"
	"google.golang.org/grpc"
)

var (
	ctx = context.Background()
)

func main() {
	conn, err :=grpc.Dial("127.0.0.1:9877", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := NewGetHelloServiceClient(conn)
	req := new(GetHelloRequest)
	req.Greeting = "ikhnaie"
	resp, err := client.Hello(ctx, req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("req: %s, resp: %s\n", req.Greeting, resp.Reply)
}
