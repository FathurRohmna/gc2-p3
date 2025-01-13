package server

import (
	"context"
	"server/model/web"
	"server/services"
	"time"

	proto "server/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BookServer struct {
	proto.UnimplementedBookServiceServer
	BookService services.IBookService
}

func NewBookServer(bookService services.IBookService) *BookServer {
	return &BookServer{
		BookService: bookService,
	}
}

func (s *BookServer) RegisterBook(ctx context.Context, req *proto.RegisterBookRequest) (*proto.BookResponse, error) {
	bookRequest := web.RegisterBookRequest{
		Title:         req.Title,
		Author:        req.Author,
		PublishedDate: req.PublishedDate,
	}

	bookResponse, err := s.BookService.RegisterBook(ctx, bookRequest)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error during book registration: %v", err)
	}

	return &proto.BookResponse{
		Id:            bookResponse.ID.String(),
		Title:         bookResponse.Title,
		Author:        bookResponse.Author,
		PublishedDate: bookResponse.PublishedDate,
		Status:        bookResponse.Status,
	}, nil
}

func (s *BookServer) BorrowBook(ctx context.Context, req *proto.BorrowBookRequest) (*proto.BorrowedBookResponse, error) {
	borrowRequest := web.BorrowBookRequest{
		BookID:       req.BookId,
		UserID:       req.UserId,
		BorrowedDate: req.BorrowedDate,
	}

	borrowResponse, err := s.BookService.BorrowBook(ctx, borrowRequest)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error during book borrowing: %v", err)
	}

	return &proto.BorrowedBookResponse{
		Id:           borrowResponse.ID,
		BookId:       borrowResponse.BookID,
		UserId:       borrowResponse.UserID,
		BorrowedDate: borrowResponse.BorrowedDate,
	}, nil
}

func (s *BookServer) ReturnBook(ctx context.Context, req *proto.ReturnBookRequest) (*proto.BorrowedBookResponse, error) {
	returnDate, err := time.Parse("2006-01-02", req.ReturnDate)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid date format: %v", err)
	}

	returnRequest := web.ReturnBookRequest{
		BookID:     req.BookId,
		UserID:     req.UserId,
		ReturnDate: returnDate.Format("2006-01-02"),
	}

	returnResponse, err := s.BookService.ReturnBook(ctx, returnRequest)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error during book return: %v", err)
	}

	return &proto.BorrowedBookResponse{
		Id:           returnResponse.ID,
		BookId:       returnResponse.BookID,
		UserId:       returnResponse.UserID,
		BorrowedDate: returnResponse.BorrowedDate,
		ReturnDate:   returnResponse.ReturnDate,
	}, nil
}
