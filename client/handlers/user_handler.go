package handlers

import (
	"client/model/web"
	"client/service"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler() (*AuthHandler, error) {
	authService, err := service.NewAuthService()
	if err != nil {
		return nil, fmt.Errorf("could not initialize AuthService: %w", err)
	}
	return &AuthHandler{authService: authService}, nil
}

// Register handles the user registration HTTP request
// @Summary Register a new user
// @Description Register a new user with username and password
// @Tags user
// @Accept json
// @Produce json
// @Param request body web.RegisterUserRequest true "Register User Request"
// @Success 200 {object} pb.RegisterResponse "User registered successfully"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /users/register [post]
func (h *AuthHandler) Register(ctx echo.Context) error {
	var req web.RegisterUserRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	resp, err := h.authService.RegisterUser(req.Username, req.Password)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, resp)
}

// Login handles the user login HTTP request
// @Summary Login an existing user
// @Description Login an existing user using username and password
// @Tags user
// @Accept json
// @Produce json
// @Param request body web.LoginUserRequest true "Login User Request"
// @Success 200 {object} pb.LoginResponse "Login successful"
// @Failure 401 {object} map[string]string "Invalid credentials"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /users/login [post]
func (h *AuthHandler) Login(ctx echo.Context) error {
	var req web.LoginUserRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	fmt.Println(req)
	resp, err := h.authService.LoginUser(req.Username, req.Password)
	if err != nil {
		if err.Error() == "invalid credentials" {
			return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, resp)
}
