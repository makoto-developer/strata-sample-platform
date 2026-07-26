package handler

import (
	"context"

	catalogv1 "github.com/makoto-developer/strata-sample-platform/proto/catalog/v1"
	"github.com/makoto-developer/strata-sample-platform/services/catalog/internal/config"
	"github.com/makoto-developer/strata-sample-platform/services/catalog/internal/searchclient"
)

// Server は CatalogService の gRPC ハンドラ。
type Server struct {
	catalogv1.UnimplementedCatalogServiceServer
	search *searchclient.Client
}

func New(cfg config.Config) *Server {
	return &Server{search: searchclient.New(cfg.SearchURL)}
}

func (s *Server) Serve(addr string) error { return nil }

func (s *Server) GetProduct(ctx context.Context, req *catalogv1.GetProductRequest) (*catalogv1.GetProductReply, error) {
	return &catalogv1.GetProductReply{Id: req.Id}, nil
}

func (s *Server) ListProducts(ctx context.Context, req *catalogv1.ListProductsRequest) (*catalogv1.ListProductsReply, error) {
	return &catalogv1.ListProductsReply{}, nil
}

// SearchProducts は検索サービス(TS / REST)に委譲する。
func (s *Server) SearchProducts(ctx context.Context, req *catalogv1.SearchProductsRequest) (*catalogv1.SearchProductsReply, error) {
	if err := s.search.Query(ctx, req.Id); err != nil {
		return nil, err
	}
	return &catalogv1.SearchProductsReply{}, nil
}

func (s *Server) UpsertProduct(ctx context.Context, req *catalogv1.UpsertProductRequest) (*catalogv1.UpsertProductReply, error) {
	if err := s.search.Index(ctx, req.Id); err != nil {
		return nil, err
	}
	return &catalogv1.UpsertProductReply{Id: req.Id}, nil
}
