package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/cylixlee/protobuf-playground/internal/proto"
	"google.golang.org/grpc"
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

func main() {
	port := os.Getenv("PP_SERVER_PORT")
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	s := grpc.NewServer()
	proto.RegisterHelloServer(s, server{})

	if err := s.Serve(l); err != nil {
		log.Fatalln(err)
	}
}
