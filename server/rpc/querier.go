package rpc

import (
	"context"
	"google.golang.org/grpc"
)

var queriers = make(map[string]any)
var hosts = make(map[string]string)

type Stringable interface {
	String() string
}

type QuerierConstructor[T any] func(cc *grpc.ClientConn) Querier[T]

func RegisterQuerier[T any](querier QuerierConstructor[T]) {
	res := interface{}(new(T))

	queriers[res.(Stringable).String()] = querier
}

func RegisterHost[T any](host string) {
	res := interface{}(new(T))

	hosts[res.(Stringable).String()] = host
}

func NewQuerier[T any](cc *grpc.ClientConn) Querier[T] {
	res := interface{}(new(T))

	queriers[res.(Stringable).String()].(QuerierConstructor[T])(cc)
}

func ConnectNewQuerier[T any](ctx context.Context) (Querier[T], error) {
	res := interface{}(new(T))

	cc, err := grpc.DialContext(ctx, hosts[res.(Stringable).String()])
	if err != nil {
		return nil, err
	}

	return queriers[res.(Stringable).String()].(QuerierConstructor[T])(cc), nil
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
