package main

import (
	"log"
	"net"
	"os"
	"server/app"
	"server/middleware"
	proto "server/pb"
	"server/repository"
	"server/server"
	"server/services"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "50051"
	}

	db := app.NewDB()

	userRepository := repository.NewUserRepository()
	userService := services.NewUserService(userRepository, db)
	userServerController := server.NewUserServer(userService)

	bookRepository := repository.NewBookRepository()
	bookService := services.NewBookService(bookRepository, db)
	bookServerController := server.NewBookServer(bookService)

	userServer := grpc.NewServer()
	bookServer := grpc.NewServer(grpc.UnaryInterceptor(middleware.AuthMiddleware()))

	proto.RegisterUserServiceServer(userServer, userServerController)
	proto.RegisterBookServiceServer(bookServer, bookServerController)

	lisUser, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	go func() {
		log.Printf("User Server listening on :50051")
		if err := userServer.Serve(lisUser); err != nil {
			log.Fatalf("Failed to serve user server: %v", err)
		}
	}()

	lisBook, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen on port 50052: %v", err)
	}

	log.Printf("Book Server listening on :50052")
	if err := bookServer.Serve(lisBook); err != nil {
		log.Fatalf("Failed to serve book server: %v", err)
	}
}
