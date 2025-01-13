package web

import "github.com/google/uuid"

type UserResponse struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}

type RegisterUserRequest struct {
	Username string `json:"username" validate:"required" example:"johndoe"`
	Password string `json:"password" validate:"required,min=8" example:"example_password"`
}

type LoginUserRequest struct {
	Username string `json:"username" validate:"required" example:"john.doe"`
	Password string `json:"password" validate:"required,min=8" example:"example_password"`
}
