package handler

import (
	"context"

	notificationv1 "github.com/makoto-developer/strata-sample-platform/proto/notification/v1"
	"github.com/makoto-developer/strata-sample-platform/services/notification/internal/mail"
)

// Server は NotificationService の gRPC ハンドラ。
type Server struct {
	notificationv1.UnimplementedNotificationServiceServer
	sender *mail.Sender
}

func New() *Server { return &Server{sender: mail.New()} }

func (s *Server) Serve(addr string) error { return nil }

func (s *Server) SendEmail(ctx context.Context, req *notificationv1.SendEmailRequest) (*notificationv1.SendEmailReply, error) {
	if err := s.sender.Send(ctx, req.Id); err != nil {
		_ = s.sender.NotifyOps(ctx, "mail failed: "+req.Id)
		return nil, err
	}
	return &notificationv1.SendEmailReply{Id: req.Id, Status: "sent"}, nil
}

func (s *Server) SendPush(ctx context.Context, req *notificationv1.SendPushRequest) (*notificationv1.SendPushReply, error) {
	return &notificationv1.SendPushReply{Id: req.Id, Status: "sent"}, nil
}

func (s *Server) ListNotifications(ctx context.Context, req *notificationv1.ListNotificationsRequest) (*notificationv1.ListNotificationsReply, error) {
	return &notificationv1.ListNotificationsReply{}, nil
}
