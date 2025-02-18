package main

import (
	"context"
	"log"
	"log/slog"
	"net"

	grpcstreaming "github.com/Emircaan/grpc-redis"
	"github.com/Emircaan/grpc-redis/notificationservice"
	"github.com/Emircaan/grpc-redis/notificationservice/notificationproto"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", notificationservice.Adrdress)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	redisClient := grpcstreaming.NewRedisClient(context.Background())
	handler := notificationservice.NewHandler(redisClient)
	notificationproto.RegisterNotificationServiceServer(grpcServer, handler)

	slog.Info("Server is running on port: ", notificationservice.Adrdress)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
