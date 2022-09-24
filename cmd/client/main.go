package main

import (
	"context"
	"log"

	pb "github.com/gdguesser/grpc-server/grpc"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewGrpcClient(conn)

	req := &pb.HelloRequest{
		Name: "Gabriel",
	}

	res, err := client.SayHello(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(res)
}
