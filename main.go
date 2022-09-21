package main

import (
	"log"
	"net"

	ggrpc "github.com/gdguesser/grpc-server/grpc"
	"google.golang.org/grpc"
)

func main() {
	println("gRPC server in Go")

	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	ggrpc.RegisterGrpcServer(s, &ggrpc.Server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
