package tcc

import (
	"context"
	"github.com/liujunren93/tcc/proto"
	"google.golang.org/grpc"
)

type Tcc interface {
	Try(ctx context.Context, in *proto.TccReq, opts ...grpc.CallOption) (*proto.TccRes, error)
	Confirm(ctx context.Context, in *proto.TccReq, opts ...grpc.CallOption) (*proto.TccRes, error)
	Cancel(ctx context.Context, in *proto.TccReq, opts ...grpc.CallOption) (*proto.TccRes, error)
}



