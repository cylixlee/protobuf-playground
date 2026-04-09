package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/cylixlee/protobuf-playground/internal/proto"
	consul "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintf(os.Stderr, "Usage: client-once <request_content>")
		return
	}

	consulClient, err := consul.NewClient(consul.DefaultConfig())
	if err != nil {
		log.Fatalln(err)
	}
	serviceMap, err := consulClient.Agent().ServicesWithFilter("Service==`hello`")
	if err != nil {
		log.Fatalln(err)
	}

	var addr string
	for k, v := range serviceMap {
		fmt.Println("Found service: ", k)
		addr = fmt.Sprintf("%s:%d", v.Address, v.Port)
	}
	fmt.Println("Choose", addr)

	c, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
