package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/cylixlee/protobuf-playground/internal/graceful"
	"github.com/cylixlee/protobuf-playground/internal/impl"
	"github.com/cylixlee/protobuf-playground/internal/proto"
	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("PP_SERVER_PORT")
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	s := grpc.NewServer()
	proto.RegisterHelloServer(s, impl.SimpleServer{})
	proto.RegisterRateLimitedHelloServer(s, impl.NewRateLimitedServer())

	graceful.New(s.Serve).
		AddParams(l).
		Defer(s.GracefulStop).
		Defer(fmt.Println, "Gracefully shutting down...").
		Run()
}
