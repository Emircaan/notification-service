package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/Emircaan/grpc-redis/notificationservice"
	"github.com/Emircaan/grpc-redis/notificationservice/notificationproto"
)

func main() {
	client, err := notificationservice.NewClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	stream, err := client.GetNotifications(ctx, &notificationproto.NotificationRequest{
		UserId: "123",
	})
	if err != nil {
		panic(err)
	}

	for {
		notification, err := stream.Recv()
		if err == io.EOF {
			break

		}
		if err != nil {
			log.Fatalf("Failed to receive a notification: %v", err)
		}
		b, err := json.MarshalIndent(notification, "", "/t")

		if err != nil {
			log.Fatalf("Failed to marshal notification: %v", err)
		}
		fmt.Println(string(b))

	}
}
