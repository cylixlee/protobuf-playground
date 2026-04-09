package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/cylixlee/protobuf-playground/internal/graceful"
	"github.com/cylixlee/protobuf-playground/internal/impl"
	"github.com/cylixlee/protobuf-playground/internal/proto"
	consul "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthproto "google.golang.org/grpc/health/grpc_health_v1"
)

const (
	serviceName = "hello"
)

var (
	outboundIP       net.IP
	outboundIPString string
	port             int
	portString       string
)

func init() {
	// Init outbound IP (net.IP and string)
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	outboundIP = conn.LocalAddr().(*net.UDPAddr).IP
	segments := make([]string, 4)
	for _, segment := range outboundIP {
		segments = append(segments, strconv.Itoa(int(segment)))
	}
	outboundIPString = strings.Join(segments, ".")

	// Init service port (int and string)
	portString = os.Getenv("PP_SERVER_PORT")
	port, err = strconv.Atoi(portString)
	if err != nil {
		panic(err)
	}
}

func main() {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	s := grpc.NewServer(
		grpc.UnaryInterceptor(logInterceptor),
	)
	proto.RegisterHelloServer(s, impl.SimpleServer{})
	proto.RegisterRateLimitedHelloServer(s, impl.NewRateLimitedServer())
	healthproto.RegisterHealthServer(s, health.NewServer())

	consulClient, err := consul.NewClient(consul.DefaultConfig())
	if err != nil {
		log.Fatalln(err)
	}

	healthCheck := &consul.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", outboundIP, port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}
	consulReg := &consul.AgentServiceRegistration{
		ID:      fmt.Sprintf("%s-%s-%d", serviceName, outboundIPString, port),
		Name:    serviceName,
		Address: outboundIPString,
		Port:    port,
		Check:   healthCheck,
	}
	consulClient.Agent().ServiceRegister(consulReg)

	graceful.New(s.Serve).
		AddParams(l).
		Defer(s.GracefulStop).
		Defer(fmt.Println, "Gracefully shutting down...").
		Run()
}

func logInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	fmt.Println(info.FullMethod)
	return handler(ctx, req)
}
