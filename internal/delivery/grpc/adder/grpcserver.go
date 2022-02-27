package adder

import (
	"github.com/JIeeiroSst/store/internal/delivery/grpc"
)

type GRPCServer struct {
	grpc *grpc.Grpc
}

func NewGRPCServer(grpc *grpc.Grpc) *GRPCServer {
	return &GRPCServer{
		grpc: grpc,
	}
}

