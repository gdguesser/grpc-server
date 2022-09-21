package grpc

import (
	context "context"
	"fmt"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  string `json:"body"`
}

type Server struct {
	UnimplementedGrpcServer
}

func (s *Server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	return &HelloReply{Message: "Hello, " + in.GetName()}, nil
}

func (s *Server) Sum(ctx context.Context, in *SumRequest) (*SumReply, error) {
	return &SumReply{Result: in.Num + in.Num2}, nil
}

func (s *Server) Health(ctx context.Context, in *HealthRequest) (*HealthReply, error) {
	return &HealthReply{Message: "UP"}, nil
}

func (*Server) AddTodo(ctx context.Context, in *TodoRequest) (*TodoReply, error) {
	todo := &Todo{}
	todos := []Todo{}

	todo.ID = len(todos) + 1

	todos = append(todos, *todo)
	fmt.Println(todos)

	return &TodoReply{Title: in.Title, Body: in.Body}, nil
}
