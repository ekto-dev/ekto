# Ekto

Ekto is a suite of tools for using protobuf "for everything" – gRPC, HTTP,
and message queue handlers. It supplements the existing `grpc-go` and `grpc-gateway`
plugins to simplify the setup of gRPC and gateway-based HTTP services.

It also exposes a way to handle cloudevents messages via an RPC service implementation,
enabling you to use protobufs to define your event schema and handlers.

## Usage

_A full example can be found in the `example` folder_.

Using protobuf, you can declare `(ekto.dev).mq.handles` to specify that an RPC method
handles a given cloudevents message. You can also use the `google.api.http` option
to have Ekto generate an HTTP gateway for your service.

```protobuf
syntax = "proto3";
option go_package = "github.com/ekto-dev/ekto/example/proto;example";

import "ekto/ekto.proto";
import "google/api/annotations.proto";
import "google/protobuf/empty.proto";

message UserCreated {
  string id = 1;
  string name = 2;
}

message NotificationRequest {
  string user_id = 1;
  string message = 2;
}

message NotificationResponse {
  string user_id = 1;
  string message = 2;
}

service NotificationService {
  rpc HandleUserCreated(UserCreated) returns (google.protobuf.Empty) {
    option (ekto.dev).mq.handles = "user.created";
  }

  rpc NotifyUser(NotificationRequest) returns (NotificationResponse) {
    option (google.api.http) = {
      post: "/v1/notify"
      body: "*"
    };
  }
}

```

This declares a service that will handle a `user.created` event from a cloudevents
message queue. Ekto will decode the cloudevent message and pass its payload to the
handler.

To run the service in Go, you should use the generated code and pass a cloudevent client (full code in the `example` folder):

```go
package main

import (
	"context"

	"github.com/Shopify/sarama"
	"github.com/cloudevents/sdk-go/protocol/kafka_sarama/v2"
	cloudeventsv2 "github.com/cloudevents/sdk-go/v2"

	"github.com/ekto-dev/ekto/example/proto"
	"github.com/ekto-dev/ekto/example/service"
)

func main() {
	proxy := example.NewNotificationServiceMQProxy()
	proxy.Register(&service.NotificationService{})
	saramaConfig := sarama.NewConfig()
	saramaConfig.Version = sarama.V2_0_0_0
	consumer, err := kafka_sarama.NewConsumer([]string{"localhost:9092"}, saramaConfig, "example", "events")
	if err != nil {
		panic(err)
	}
	cloudeventsClient, err := cloudeventsv2.NewClient(consumer, cloudeventsv2.WithTimeNow(), cloudeventsv2.WithUUIDs())
	if err != nil {
		panic(err)
	}

	proxy.Run(context.Background(), cloudeventsClient)
}
```

Ekto also provides an opinionated setup for running your gRPC and HTTP Gateway
service, using the default `net/http` setup:

```go
example.StartNotificationServiceServer(
	context.Background(),
	":50051",
	&service.NotificationService{},
	server.WithGateway(":8080"),
)
```

Note that the gateway will only be started if the `google.api.http` option is
used in your protobuf file.

### Querying other gRPC services

If you want a monolith-like developer experience, Ekto provides a `Querier` which currently
supports `Find` (for now), allowing you to get a resource by ID from another service.

This makes it possible to write code like this:
```go
user, err := rpc.Find[usersgen.User](ctx, request.UserId)
```

Without having to ever configure the gRPC connection to the users service. Instead, the generated
code will map a given message type to the right methods. To support this, you need to add some
protobuf options to your service being queried. Namely, `(ekto.msg).queryable = true` on the "model"
that's being queried, and `(ekto.dev).querier.method` on the method:
```protobuf
syntax = "proto3";

// GetUserRequest and anything else needed

message User {
  option (ekto.msg).queryable = true;
  string id = 1;
  string name = 2;
}

service UserService {
  rpc GetUser(GetUserRequest) returns (User) {
    option (ekto.dev).querier.method = FIND;
  }
}
```

## Limitations

- If you specify the `google.api.http` option, you **must** include the `grpc-gateway`
  protobuf plugin, as Ekto relies on the generated gateway code to in _its_ generated output.
