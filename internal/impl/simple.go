package impl

import (
	"context"
	"fmt"

	"github.com/cylixlee/protobuf-playground/internal/proto"
)

type SimpleServer struct {
	proto.UnimplementedHelloServer
}

func (SimpleServer) Hello(_ context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	resp := &proto.HelloResponse{
		Reply: fmt.Sprintf("Hello %s!", req.Name),
	}
	return resp, nil
}
