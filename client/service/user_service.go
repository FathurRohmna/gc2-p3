package service

import (
	"context"
	"fmt"
	"time"

	"client/grpc"
	pb "client/pb"
)

type AuthService struct {
	grpcClient pb.UserServiceClient
}

func NewAuthService() (*AuthService, error) {
	client, err := grpc.NewGRPCClient()
	if err != nil {
		return nil, fmt.Errorf("could not initialize gRPC client: %w", err)
	}
	return &AuthService{grpcClient: client}, nil
}

func (s *AuthService) RegisterUser(username, password string) (*pb.RegisterResponse, error) {
	req := &pb.RegisterRequest{
		Username: username,
		Password: password,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return s.grpcClient.Register(ctx, req)
}

func (s *AuthService) LoginUser(username, password string) (*pb.LoginResponse, error) {
	req := &pb.LoginRequest{
		Username: username,
		Password: password,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return s.grpcClient.Login(ctx, req)
}
