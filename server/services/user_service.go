package services

import (
	"context"
	"server/model/web"
)

type IUserService interface {
	Login(ctx context.Context, student web.LoginUserRequest) string
	Register(ctx context.Context, student web.RegisterUserRequest) web.UserResponse
}
