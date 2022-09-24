package main

import (
	"context"
	"log"
	"net"

	ggrpc "github.com/gdguesser/grpc-server/grpc"
	"google.golang.org/grpc"
)

type Server struct {
	ggrpc.UnimplementedGrpcServer
}

func (s *Server) SayHello(ctx context.Context, in *ggrpc.HelloRequest) (*ggrpc.HelloReply, error) {
	return &ggrpc.HelloReply{Message: "Hello, " + in.GetName()}, nil
}

func main() {
	println("gRPC server in Go")

	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	ggrpc.RegisterGrpcServer(s, &Server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
