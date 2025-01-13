package main

import (
	"client/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	_ "client/docs"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Exoplanet API
// @version 1.0
// @description This is the API for planet.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	authHandler, err := handlers.NewAuthHandler()
	if err != nil {
		log.Fatalf("Failed to initialize auth handler: %v", err)
	}

	bookHandler, err := handlers.NewBookHandler()
	if err != nil {
		log.Fatalf("Failed to initialize auth handler: %v", err)
	}

	usersRouter := e.Group("/users")
	usersRouter.POST("/register", authHandler.Register)
	usersRouter.POST("/login", authHandler.Login)

	bookRoutes := e.Group("/books")
	bookRoutes.POST("/register", bookHandler.RegisterBook)
	bookRoutes.POST("/borrow", bookHandler.BorrowBook)
	bookRoutes.POST("/return", bookHandler.ReturnBook)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	go func() {
		if err := e.Start(":" + port); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Shutting down the server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutting down server...")
	if err := e.Shutdown(context.Background()); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

}
