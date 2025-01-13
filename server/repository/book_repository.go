package repository

import (
	"context"
	"server/model/domain"

	"gorm.io/gorm"
)

type IBookRepository interface {
	Save(ctx context.Context, tx *gorm.DB, book domain.Book) (domain.Book, error)
	FindBookByIDAndStatus(ctx context.Context, tx *gorm.DB, bookID string, status string) (domain.Book, error)
	CreateBorrowedBook(ctx context.Context, tx *gorm.DB, borrowedBook domain.BorrowedBook) (domain.BorrowedBook, error)
	FindBorrowedBook(ctx context.Context, tx *gorm.DB, bookID, userID string) (domain.BorrowedBook, error)
	UpdateBookStatus(ctx context.Context, tx *gorm.DB, book domain.Book) (domain.Book, error)
	UpdateBorrowedBookReturnDate(ctx context.Context, tx *gorm.DB, borrowedBook domain.BorrowedBook) error
}
