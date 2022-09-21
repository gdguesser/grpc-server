package grpc

import context "context"

type Server struct {
	UnimplementedGrpcServer
}

func (s *Server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	return &HelloReply{Message: "Hello, " + in.GetName()}, nil
}

func (s *Server) Sum(ctx context.Context, in *SumRequest) (*SumReply, error) {
	return &SumReply{Result: in.Num+in.Num2}, nil
}
