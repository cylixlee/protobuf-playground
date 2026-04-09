package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"

	"github.com/cylixlee/protobuf-playground/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	proto.UnimplementedHelloServer
}

func (server) Hello(_ context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	resp := &proto.HelloResponse{
		Reply: fmt.Sprintf("Hello %s!", req.Name),
	}
	return resp, nil
}

type rateLimitedServer struct {
	proto.UnimplementedRateLimitedHelloServer
	visitors sync.Map
}

func newRateLimitedServer() *rateLimitedServer {
	return new(rateLimitedServer)
}

func (r *rateLimitedServer) Hello(_ context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	if _, exist := r.visitors.Load(req.Name); exist {
		return nil, status.Error(codes.AlreadyExists, fmt.Sprintf("User %s already exists", req.Name))
	}
	r.visitors.Store(req.Name, struct{}{})
	return &proto.HelloResponse{Reply: fmt.Sprintf("Hello %s", req.Name)}, nil
}

func main() {
	port := os.Getenv("PP_SERVER_PORT")
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	s := grpc.NewServer()
	proto.RegisterHelloServer(s, server{})
	proto.RegisterRateLimitedHelloServer(s, newRateLimitedServer())

	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt)

	go func() {
		if err := s.Serve(l); err != nil {
			log.Fatalln(err)
		}
	}()

	<-interruptChan
	fmt.Println("Graceful shutting down...")
	s.GracefulStop()
}
