package rpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var queriers = make(map[string]any)
var hosts = make(map[string]string)

type QueryableMessage interface {
	EktoQualifiedName() string
}

func getName[T QueryableMessage]() string {
	msg := *new(T)
	return msg.EktoQualifiedName()
}

type QuerierConstructor[T QueryableMessage] func(cc *grpc.ClientConn) Querier[T]

func RegisterQuerier[T QueryableMessage](querier QuerierConstructor[T]) {
	queriers[getName[T]()] = querier
}

func RegisterHost[T QueryableMessage](host string) {
	hosts[getName[T]()] = host
}

func NewQuerier[T QueryableMessage](cc *grpc.ClientConn) Querier[T] {
	return queriers[getName[T]()].(QuerierConstructor[T])(cc)
}

func ConnectNewQuerier[T QueryableMessage](ctx context.Context) (Querier[T], error) {
	cc, err := grpc.DialContext(ctx, hosts[getName[T]()], grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return queriers[getName[T]()].(QuerierConstructor[T])(cc), nil
}

type Querier[T QueryableMessage] interface {
	Find(ctx context.Context, id string) (*T, error)
}

func Find[T QueryableMessage](ctx context.Context, id string) (*T, error) {
	querier, err := ConnectNewQuerier[T](ctx)
	if err != nil {
		return new(T), err
	}

	return querier.Find(ctx, id)
}
