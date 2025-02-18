package notificationservice

import (
	"fmt"

	"github.com/Emircaan/grpc-redis/notificationservice/notificationproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient() (notificationproto.NotificationServiceClient, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(Adrdress, opts...)
	if err != nil {
		return nil, fmt.Errorf("Failed to dial: %v", err)
	}
	return notificationproto.NewNotificationServiceClient(conn), nil

}
