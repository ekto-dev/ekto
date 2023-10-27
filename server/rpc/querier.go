package rpc

import (
	"context"
	"google.golang.org/grpc"
)

var queriers = make(map[any]any)
var hosts = make(map[any]string)

type QuerierConstructor[T any] func(cc *grpc.ClientConn) Querier[T]

func RegisterQuerier[T any](querier QuerierConstructor[T]) {
	res := *new(T)
	queriers[res] = querier
}

func RegisterHost[T any](host string) {
	res := *new(T)
	hosts[res] = host
}

func NewQuerier[T any](cc *grpc.ClientConn) Querier[T] {
	res := *new(T)

	return queriers[res].(QuerierConstructor[T])(cc)
}

func ConnectNewQuerier[T any](ctx context.Context) (Querier[T], error) {
	res := *new(T)

	cc, err := grpc.DialContext(ctx, hosts[res])
	if err != nil {
		return nil, err
	}

	return queriers[res].(QuerierConstructor[T])(cc), nil
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
