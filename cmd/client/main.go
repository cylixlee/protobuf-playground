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
	port := os.Getenv("PP_SERVER_PORT")
	c, err := grpc.NewClient(fmt.Sprintf("localhost:%s", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	service := proto.NewHelloClient(c)
	resp, err := service.Hello(context.Background(), &proto.HelloRequest{Name: "CYLIX"})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp.Reply)
}
