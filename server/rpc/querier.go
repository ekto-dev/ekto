package rpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var queriers = make(map[protoreflect.FullName]any)
var hosts = make(map[protoreflect.FullName]string)

func getName[T any]() protoreflect.FullName {
	res := interface{}(new(T))

	return res.(protoreflect.ProtoMessage).ProtoReflect().Descriptor().FullName()
}

type QuerierConstructor[T any] func(cc *grpc.ClientConn) Querier[T]

func RegisterQuerier[T any](querier QuerierConstructor[T]) {
	queriers[getName[T]()] = querier
}

func RegisterHost[T any](host string) {
	hosts[getName[T]()] = host
}

func NewQuerier[T any](cc *grpc.ClientConn) Querier[T] {
	return queriers[getName[T]()].(QuerierConstructor[T])(cc)
}

func ConnectNewQuerier[T any](ctx context.Context) (Querier[T], error) {
	cc, err := grpc.DialContext(ctx, hosts[getName[T]()])
	if err != nil {
		return nil, err
	}

	return queriers[getName[T]()].(QuerierConstructor[T])(cc), nil
}

type Querier[T any] interface {
	Find(ctx context.Context, id string) (*T, error)
}

func Find[T any](ctx context.Context, id string) (*T, error) {
	querier, err := ConnectNewQuerier[T](ctx)
	if err != nil {
		return nil, err
	}

	return querier.Find(ctx, id)
}
