package handler

import (
	"context"
	"net"

	userv1 "github.com/makoto-developer/strata-sample-platform/proto/user/v1"
	"github.com/makoto-developer/strata-sample-platform/services/user/internal/store"
)

// Server は UserService の gRPC ハンドラ。
type Server struct {
	userv1.UnimplementedUserServiceServer
	store *store.Store
}

func New(s *store.Store) *Server {
	return &Server{store: s}
}

func (s *Server) Serve(lis net.Listener) error {
	return nil
}

// GetUser はユーザーを 1 件返す gRPC ハンドラ。
func (s *Server) GetUser(ctx context.Context, req *userv1.GetUserRequest) (*userv1.GetUserReply, error) {
	name, err := s.store.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &userv1.GetUserReply{Id: req.Id, Status: name}, nil
}

func (s *Server) ListUsers(ctx context.Context, req *userv1.ListUsersRequest) (*userv1.ListUsersReply, error) {
	if _, err := s.store.List(ctx); err != nil {
		return nil, err
	}
	return &userv1.ListUsersReply{}, nil
}

func (s *Server) CreateUser(ctx context.Context, req *userv1.CreateUserRequest) (*userv1.CreateUserReply, error) {
	if err := s.store.Save(ctx, req.Id, req.ActorId); err != nil {
		return nil, err
	}
	return &userv1.CreateUserReply{Id: req.Id}, nil
}

func (s *Server) UpdateProfile(ctx context.Context, req *userv1.UpdateProfileRequest) (*userv1.UpdateProfileReply, error) {
	if err := s.store.Save(ctx, req.Id, req.ActorId); err != nil {
		return nil, err
	}
	return &userv1.UpdateProfileReply{Id: req.Id}, nil
}

func (s *Server) DeactivateUser(ctx context.Context, req *userv1.DeactivateUserRequest) (*userv1.DeactivateUserReply, error) {
	if err := s.store.Deactivate(ctx, req.Id); err != nil {
		return nil, err
	}
	return &userv1.DeactivateUserReply{Id: req.Id}, nil
}
