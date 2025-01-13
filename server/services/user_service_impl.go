package services

import (
	"context"
	"log"
	"os"
	"server/helper"
	"server/model/domain"
	"server/model/web"
	"server/repository"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserService struct {
	UserRepository repository.IUserRepository
	DB             *gorm.DB
}

func NewUserService(userRepository repository.IUserRepository, DB *gorm.DB) *UserService {
	return &UserService{
		UserRepository: userRepository,
		DB:             DB,
	}
}

func (service *UserService) Login(ctx context.Context, userRequest web.LoginUserRequest) (string, error) {
	tx := service.DB.Begin()

	user, err := service.UserRepository.FindByUsername(ctx, tx, userRequest.Username)
	if err != nil {
		if err.Error() == "user not found" {
			return "", status.Errorf(codes.NotFound, "Username not found")
		}
		return "", status.Errorf(codes.Internal, "Internal error: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userRequest.Password))
	if err != nil {
		return "", status.Errorf(codes.Unauthenticated, "Invalid username or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", status.Errorf(codes.Internal, "Error signing token: %v", err)
	}

	return tokenString, nil
}

func (service *UserService) Register(ctx context.Context, userRequest web.RegisterUserRequest) (web.UserResponse, error) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	_, err := service.UserRepository.FindByUsername(ctx, tx, userRequest.Username)
	if err == nil {
		return web.UserResponse{}, status.Errorf(codes.AlreadyExists, "Username already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), 10)
	if err != nil {
		return web.UserResponse{}, status.Errorf(codes.Internal, "Error hashing password: %v", err)
	}

	user := domain.User{
		Password: string(hash),
		Username: userRequest.Username,
	}
	createdUser, err := service.UserRepository.Save(ctx, tx, user)
	if err != nil {
		return web.UserResponse{}, status.Errorf(codes.Internal, "Error saving user: %v", err)
	}

	userID, err := uuid.Parse(createdUser.ID)
	if err != nil {
		log.Fatalf("Invalid user ID: %v", err)
		return web.UserResponse{}, status.Errorf(codes.Internal, "Invalid user ID format")
	}

	return web.UserResponse{
		ID:       userID,
		Username: createdUser.Username,
	}, nil
}
