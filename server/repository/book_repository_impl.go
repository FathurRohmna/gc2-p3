package repository

import (
	"context"
	"errors"

	"server/model/domain"

	"gorm.io/gorm"
)

type BookRepository struct{}

func NewBookRepository() IBookRepository {
	return &BookRepository{}
}

func (r *BookRepository) Save(ctx context.Context, tx *gorm.DB, book domain.Book) (domain.Book, error) {
	err := tx.WithContext(ctx).Save(&book).Error
	if err != nil {
		return domain.Book{}, err
	}
	return book, nil
}

func (r *BookRepository) FindBookByIDAndStatus(ctx context.Context, tx *gorm.DB, bookID string, status string) (domain.Book, error) {
	var book domain.Book
	err := tx.WithContext(ctx).Where("id = ?", bookID).First(&book).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.Book{}, errors.New("book not found or not available")
		}
		return domain.Book{}, err
	}
	return book, nil
}

func (r *BookRepository) CreateBorrowedBook(ctx context.Context, tx *gorm.DB, borrowedBook domain.BorrowedBook) (domain.BorrowedBook, error) {
	err := tx.WithContext(ctx).Create(&borrowedBook).Error
	if err != nil {
		return domain.BorrowedBook{}, err
	}

	return borrowedBook, nil
}

func (r *BookRepository) FindBorrowedBook(ctx context.Context, tx *gorm.DB, bookID, userID string) (domain.BorrowedBook, error) {
	var borrowedBook domain.BorrowedBook
	err := tx.WithContext(ctx).Where("book_id = ? AND user_id = ? AND return_date IS NULL", bookID, userID).First(&borrowedBook).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.BorrowedBook{}, errors.New("no active borrowing record found")
		}
		return domain.BorrowedBook{}, err
	}
	return borrowedBook, nil
}

func (r *BookRepository) UpdateBookStatus(ctx context.Context, tx *gorm.DB, book domain.Book) (domain.Book, error) {
	err := tx.WithContext(ctx).Model(&book).Updates(map[string]interface{}{
		"status": book.Status,
	}).Error
	if err != nil {
		return domain.Book{}, err
	}
	return book, nil
}

func (r *BookRepository) UpdateBorrowedBookReturnDate(ctx context.Context, tx *gorm.DB, borrowedBook domain.BorrowedBook) error {
	err := tx.WithContext(ctx).Model(&borrowedBook).Updates(map[string]interface{}{
		"return_date": borrowedBook.ReturnDate,
	}).Error

	if err != nil {
		return err
	}

	return nil
}
