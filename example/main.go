package main

import (
	"context"
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/cloudevents/sdk-go/protocol/kafka_sarama/v2"
	cloudeventsv2 "github.com/cloudevents/sdk-go/v2"
	"github.com/ekto-dev/ekto/server"
	"golang.org/x/sync/errgroup"

	"github.com/ekto-dev/ekto/example/proto"
	"github.com/ekto-dev/ekto/example/service"
)

func main() {
	eg := errgroup.Group{}
	eg.Go(runProxy)

	eg.Go(func() error {
		fmt.Println("starting server")

		return example.StartNotificationServiceServer(
			context.Background(),
			":50051",
			&service.NotificationService{},
			server.WithGateway(":8080"),
		)
	})

	if err := eg.Wait(); err != nil {
		panic(err)
	}
}

func runProxy() error {
	proxy := example.NewNotificationServiceMQProxy()
	proxy.Register(&service.NotificationService{})
	saramaConfig := sarama.NewConfig()
	saramaConfig.Version = sarama.V2_0_0_0
	consumer, err := kafka_sarama.NewConsumer([]string{"localhost:9092"}, saramaConfig, "example", "events")
	if err != nil {
		return err
	}
	cloudeventsClient, err := cloudeventsv2.NewClient(consumer, cloudeventsv2.WithTimeNow(), cloudeventsv2.WithUUIDs())
	if err != nil {
		return err
	}

	fmt.Println("starting proxy")
	return proxy.Run(context.Background(), cloudeventsClient)
}
