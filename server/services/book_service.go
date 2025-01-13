package services

import (
	"context"
	"server/model/web"
)

type IBookService interface {
	RegisterBook(ctx context.Context, req web.RegisterBookRequest) (web.BookResponse, error)
	BorrowBook(ctx context.Context, req web.BorrowBookRequest) (web.BorrowedBookResponse, error)
	ReturnBook(ctx context.Context, req web.ReturnBookRequest) (web.BorrowedBookResponse, error)
}
