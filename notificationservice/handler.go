package notificationservice

import (
	"fmt"
	"time"

	"github.com/Emircaan/grpc-redis/notificationservice/notificationproto"
	"github.com/redis/go-redis/v9"
)

var _ notificationproto.NotificationServiceServer = (*Handler)(nil)

type Handler struct {
	notificationproto.UnimplementedNotificationServiceServer
	redisClient *redis.Client
}

func NewHandler(redisClient *redis.Client) *Handler {
	return &Handler{redisClient: redisClient}
}

func (h *Handler) GetNotifications(req *notificationproto.NotificationRequest, stream notificationproto.NotificationService_GetNotificationsServer) error {
	pubsub := h.redisClient.Subscribe(stream.Context(), fmt.Sprintf("notifications/%s", req.GetUserId()))
	for {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		case msg := <-pubsub.Channel():
			if err := stream.Send(
				&notificationproto.Notification{
					UserId:    req.GetUserId(),
					Content:   fmt.Sprintf("Notification for user %s: %s", req.GetUserId(), msg.Payload),
					CreatedAt: time.Now().Unix(),
				},
			); err != nil {
				return fmt.Errorf("could not send notification: %v", err)
			}
		}
	}
}
