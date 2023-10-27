package rpc

import (
	"context"
	"google.golang.org/grpc"
)

var queriers = make(map[any]any)

type QuerierConstructor[T any] func(cc *grpc.ClientConn) Querier[T]

func RegisterQuerier[T any](querier QuerierConstructor[T]) {
	res := *new(T)
	queriers[res] = querier
}

func NewQuerier[T any](cc *grpc.ClientConn) Querier[T] {
	res := *new(T)
	return queriers[res].(QuerierConstructor[T])(cc)
}

type Querier[T any] interface {
	Find(ctx context.Context, id string) (*T, error)
}

func Find[T any](ctx context.Context, id string) (*T, error) {
	querier := NewQuerier[T]()
	return querier.Find(ctx, id)
}
