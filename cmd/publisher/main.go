package main

import (
	"context"
	"fmt"
	"log"
	"time"

	grpcstreaming "github.com/Emircaan/grpc-redis"
)

func main() {
	ctx := context.Background()
	redisClient := grpcstreaming.NewRedisClient(ctx)
	channelName := fmt.Sprintf("notifications/%s", "123")
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ctx.Done():
			log.Println("Publisher is shutting down")
		case t := <-ticker.C:
			if cmd := redisClient.Publish(ctx, channelName, fmt.Sprintf("New notification %s", t.String())); cmd.Err() != nil {
				panic(cmd.Err())
			}
		}
	}

}
