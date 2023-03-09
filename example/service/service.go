package service

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	example "github.com/ekto-dev/ekto/example/proto"
)

type NotificationService struct {
	*example.UnimplementedNotificationServiceServer
}

func (n *NotificationService) HandleUserCreated(ctx context.Context, created *example.UserCreated) (*emptypb.Empty, error) {
	// do something with the event
	return &emptypb.Empty{}, nil
}

func (n *NotificationService) NotifyUser(ctx context.Context, request *example.NotificationRequest) (*example.NotificationResponse, error) {
	// send a notification
	return &example.NotificationResponse{}, nil
}
