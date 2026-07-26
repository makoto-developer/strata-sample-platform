package handler

import (
	"context"

	paymentv1 "github.com/makoto-developer/strata-sample-platform/proto/payment/v1"
	"github.com/makoto-developer/strata-sample-platform/services/payment/internal/gateway"
)

// Server は PaymentService の gRPC ハンドラ。
type Server struct {
	paymentv1.UnimplementedPaymentServiceServer
	provider *gateway.Provider
}

func New() *Server { return &Server{provider: gateway.New()} }

func (s *Server) Serve(addr string) error { return nil }

func (s *Server) CreatePaymentIntent(ctx context.Context, req *paymentv1.CreatePaymentIntentRequest) (*paymentv1.CreatePaymentIntentReply, error) {
	return &paymentv1.CreatePaymentIntentReply{Id: req.Id, Status: "requires_confirmation"}, nil
}

// ConfirmPayment は外部決済プロバイダに課金を送る。
func (s *Server) ConfirmPayment(ctx context.Context, req *paymentv1.ConfirmPaymentRequest) (*paymentv1.ConfirmPaymentReply, error) {
	if err := s.provider.Charge(ctx, req.Id); err != nil {
		return nil, err
	}
	return &paymentv1.ConfirmPaymentReply{Id: req.Id, Status: "paid"}, nil
}

func (s *Server) RefundPayment(ctx context.Context, req *paymentv1.RefundPaymentRequest) (*paymentv1.RefundPaymentReply, error) {
	if err := s.provider.Refund(ctx, req.Id); err != nil {
		return nil, err
	}
	return &paymentv1.RefundPaymentReply{Id: req.Id, Status: "refunded"}, nil
}

func (s *Server) GetPayment(ctx context.Context, req *paymentv1.GetPaymentRequest) (*paymentv1.GetPaymentReply, error) {
	return &paymentv1.GetPaymentReply{Id: req.Id}, nil
}
