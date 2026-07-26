package handler

import (
	"context"

	orderv1 "github.com/makoto-developer/strata-sample-platform/proto/order/v1"
	"github.com/makoto-developer/strata-sample-platform/services/order/internal/usecase"
)

// Server は OrderService の gRPC ハンドラ。
type Server struct {
	orderv1.UnimplementedOrderServiceServer
	uc *usecase.OrderUsecase
}

func New(uc *usecase.OrderUsecase) *Server { return &Server{uc: uc} }

func (s *Server) Serve(addr string) error { return nil }

func (s *Server) CreateOrder(ctx context.Context, req *orderv1.CreateOrderRequest) (*orderv1.CreateOrderReply, error) {
	if err := s.uc.Create(ctx, req.Id); err != nil {
		return nil, err
	}
	return &orderv1.CreateOrderReply{Id: req.Id, Status: "created"}, nil
}

func (s *Server) GetOrder(ctx context.Context, req *orderv1.GetOrderRequest) (*orderv1.GetOrderReply, error) {
	if err := s.uc.Track(ctx, req.Id); err != nil {
		return nil, err
	}
	return &orderv1.GetOrderReply{Id: req.Id}, nil
}

func (s *Server) ListOrders(ctx context.Context, req *orderv1.ListOrdersRequest) (*orderv1.ListOrdersReply, error) {
	return &orderv1.ListOrdersReply{}, nil
}

func (s *Server) CancelOrder(ctx context.Context, req *orderv1.CancelOrderRequest) (*orderv1.CancelOrderReply, error) {
	if err := s.uc.Cancel(ctx, req.Id); err != nil {
		return nil, err
	}
	return &orderv1.CancelOrderReply{Id: req.Id, Status: "cancelled"}, nil
}
