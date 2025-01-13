package services

import (
	"context"
	"errors"
	"server/helper"
	"server/model/domain"
	"server/model/web"
	"server/repository"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type BookService struct {
	BookRepository repository.IBookRepository
	DB             *gorm.DB
}

func NewBookService(bookRepository repository.IBookRepository, DB *gorm.DB) IBookService {
	return &BookService{
		BookRepository: bookRepository,
		DB:             DB,
	}
}

func (service *BookService) RegisterBook(ctx context.Context, req web.RegisterBookRequest) (web.BookResponse, error) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	publishedDate, err := time.Parse("2006-01-02", req.PublishedDate)
	if err != nil {
		return web.BookResponse{}, status.Errorf(codes.InvalidArgument, "Invalid date format: %v", err)
	}

	book := domain.Book{
		Title:         req.Title,
		Author:        req.Author,
		PublishedDate: publishedDate,
		Status:        "Available",
	}

	_, err = service.BookRepository.Save(ctx, tx, book)
	if err != nil {
		return web.BookResponse{}, err
	}

	return web.BookResponse{}, nil
}

func (service *BookService) BorrowBook(ctx context.Context, req web.BorrowBookRequest) (web.BorrowedBookResponse, error) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	book, err := service.BookRepository.FindBookByIDAndStatus(ctx, tx, req.BookID, "Available")
	if err != nil {
		if err.Error() == "book not found or not available" {
			return web.BorrowedBookResponse{}, status.Errorf(codes.NotFound, "Book not found")
		}
		return web.BorrowedBookResponse{}, status.Errorf(codes.Internal, "Error during book borrowing: %v", err)
	}

	borrowedBook := domain.BorrowedBook{
		BookID:       req.BookID,
		UserID:       req.UserID,
		BorrowedDate: time.Now(),
	}

	_, err = service.BookRepository.CreateBorrowedBook(ctx, tx, borrowedBook)
	if err != nil {
		return web.BorrowedBookResponse{}, err
	}

	book.Status = "Borrowed"
	_, err = service.BookRepository.UpdateBookStatus(ctx, tx, book)
	if err != nil {
		return web.BorrowedBookResponse{}, err
	}

	return web.BorrowedBookResponse{}, nil
}

func (service *BookService) ReturnBook(ctx context.Context, req web.ReturnBookRequest) (web.BorrowedBookResponse, error) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	borrowedBook, err := service.BookRepository.FindBorrowedBook(ctx, tx, req.BookID, req.UserID)
	if err != nil {
		return web.BorrowedBookResponse{}, errors.New("no active borrowing record found")
	}

	returnDate, err := time.Parse("2006-01-02", req.ReturnDate)
	if err != nil {
		return web.BorrowedBookResponse{}, status.Errorf(codes.InvalidArgument, "Invalid date format: %v", err)
	}

	borrowedBook.ReturnDate = returnDate
	if err := service.BookRepository.UpdateBorrowedBookReturnDate(ctx, tx, borrowedBook); err != nil {
		return web.BorrowedBookResponse{}, err
	}

	book, err := service.BookRepository.FindBookByIDAndStatus(ctx, tx, req.BookID, "Borrowed")
	if err != nil {
		return web.BorrowedBookResponse{}, err
	}
	book.Status = "Available"
	_, err = service.BookRepository.UpdateBookStatus(ctx, tx, book)
	if err != nil {
		return web.BorrowedBookResponse{}, err
	}

	return web.BorrowedBookResponse{}, nil
}
