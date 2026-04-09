package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/cylixlee/protobuf-playground/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintf(os.Stderr, "Usage: client-once <request_content>")
		return
	}

	port := os.Getenv("PP_SERVER_PORT")
	c, err := grpc.NewClient(fmt.Sprintf("localhost:%s", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	helloService := proto.NewHelloClient(c)
	resp, err := helloService.Hello(context.Background(), &proto.HelloRequest{Name: os.Args[1]})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Println(resp.Reply)
	}

	rateLimitedHelloService := proto.NewRateLimitedHelloClient(c)
	resp, err = rateLimitedHelloService.Hello(context.Background(), &proto.HelloRequest{Name: os.Args[1]})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Println(resp.Reply)
	}
}
