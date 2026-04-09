package impl

import (
	"context"
	"fmt"
	"sync"

	"github.com/cylixlee/protobuf-playground/internal/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RateLimitedServer struct {
	proto.UnimplementedRateLimitedHelloServer
	visitors sync.Map
}

func NewRateLimitedServer() *RateLimitedServer {
	return new(RateLimitedServer)
}

func (r *RateLimitedServer) Hello(_ context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	if _, exist := r.visitors.Load(req.Name); exist {
		return nil, status.Error(codes.AlreadyExists, fmt.Sprintf("User %s already exists", req.Name))
	}
	r.visitors.Store(req.Name, struct{}{})
	return &proto.HelloResponse{Reply: fmt.Sprintf("Hello %s", req.Name)}, nil
}
