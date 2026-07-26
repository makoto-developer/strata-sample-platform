package handler

import (
	"context"
	"sync"

	inventoryv1 "github.com/makoto-developer/strata-sample-platform/proto/inventory/v1"
)

// Server は InventoryService の gRPC ハンドラ。在庫の引当・解放・確定を扱う。
type Server struct {
	inventoryv1.UnimplementedInventoryServiceServer
	mu    sync.Mutex
	stock map[string]int
}

func New() *Server { return &Server{stock: map[string]int{}} }

func (s *Server) Serve(addr string) error { return nil }

func (s *Server) GetStock(ctx context.Context, req *inventoryv1.GetStockRequest) (*inventoryv1.GetStockReply, error) {
	return &inventoryv1.GetStockReply{Id: req.Id}, nil
}

func (s *Server) ReserveStock(ctx context.Context, req *inventoryv1.ReserveStockRequest) (*inventoryv1.ReserveStockReply, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.stock[req.Id]--
	return &inventoryv1.ReserveStockReply{Id: req.Id, Status: "reserved"}, nil
}

func (s *Server) ReleaseStock(ctx context.Context, req *inventoryv1.ReleaseStockRequest) (*inventoryv1.ReleaseStockReply, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.stock[req.Id]++
	return &inventoryv1.ReleaseStockReply{Id: req.Id, Status: "released"}, nil
}

func (s *Server) CommitStock(ctx context.Context, req *inventoryv1.CommitStockRequest) (*inventoryv1.CommitStockReply, error) {
	return &inventoryv1.CommitStockReply{Id: req.Id, Status: "committed"}, nil
}
