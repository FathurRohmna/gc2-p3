package service

import (
	"context"
	"fmt"
	"time"

	"client/grpc"
	pb "client/pb"
)

type BookService struct {
	grpcClient pb.BookServiceClient
}

func NewBookService() (*BookService, error) {
	client, err := grpc.NewGRPCBookClient()
	if err != nil {
		return nil, fmt.Errorf("could not initialize gRPC client: %w", err)
	}
	return &BookService{grpcClient: client}, nil
}

func (s *BookService) RegisterBook(title, author, publishedDate string, jwtToken string) (*pb.BookResponse, error) {
	req := &pb.RegisterBookRequest{
		Title:         title,
		Author:        author,
		PublishedDate: publishedDate,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ctx = grpc.AddJWTToContext(ctx, jwtToken)

	return s.grpcClient.RegisterBook(ctx, req)
}

func (s *BookService) BorrowBook(bookID, userID, borrowedDate string, jwtToken string) (*pb.BorrowedBookResponse, error) {
	req := &pb.BorrowBookRequest{
		BookId:       bookID,
		UserId:       userID,
		BorrowedDate: borrowedDate,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ctx = grpc.AddJWTToContext(ctx, jwtToken)

	return s.grpcClient.BorrowBook(ctx, req)
}

func (s *BookService) ReturnBook(bookID, userID, returnDate string, jwtToken string) (*pb.BorrowedBookResponse, error) {
	req := &pb.ReturnBookRequest{
		BookId:     bookID,
		UserId:     userID,
		ReturnDate: returnDate,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ctx = grpc.AddJWTToContext(ctx, jwtToken)

	return s.grpcClient.ReturnBook(ctx, req)
}
