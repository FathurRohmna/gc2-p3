package server

import (
	"context"
	"server/model/web"
	"server/services"

	proto "server/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	proto.UnimplementedUserServiceServer
	UserService *services.UserService
}

func NewUserServer(userService *services.UserService) *UserServer {
	return &UserServer{
		UserService: userService,
	}
}

func (s *UserServer) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	userRequest := web.RegisterUserRequest{
		Username: req.Username,
		Password: req.Password,
	}

	_, err := s.UserService.Register(ctx, userRequest)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error during registration: %v", err)
	}

	return &proto.RegisterResponse{
		Message: "Created user",
	}, nil
}

func (s *UserServer) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	loginRequest := web.LoginUserRequest{
		Username: req.Username,
		Password: req.Password,
	}

	token, err := s.UserService.Login(ctx, loginRequest)
	if err != nil {
		if err.Error() == "invalid credentials" {
			return nil, status.Errorf(codes.Unauthenticated, "Invalid username or password")
		}
		return nil, status.Errorf(codes.Internal, "Error during login: %v", err)
	}

	return &proto.LoginResponse{
		Token: token,
	}, nil
}
