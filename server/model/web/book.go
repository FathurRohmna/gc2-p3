package web

import (
	"github.com/google/uuid"
)

type RegisterBookRequest struct {
	Title         string `json:"title" validate:"required"`
	Author        string `json:"author" validate:"required"`
	PublishedDate string `json:"published_date"`
}

type BookResponse struct {
	ID            uuid.UUID `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	PublishedDate string    `json:"published_date"`
	Status        string    `json:"status"`
}

type BorrowBookRequest struct {
	BookID       string `json:"book_id" validate:"required"`
	UserID       string `json:"user_id" validate:"required"`
	BorrowedDate string `json:"borrowed_date"`
}

type BorrowedBookResponse struct {
	ID           string `json:"id"`
	BookID       string `json:"book_id"`
	UserID       string `json:"user_id"`
	BorrowedDate string `json:"borrowed_date"`
	ReturnDate   string `json:"return_date,omitempty"`
}

type ReturnBookRequest struct {
	BookID     string `json:"book_id" validate:"required"`
	UserID     string `json:"user_id" validate:"required"`
	ReturnDate string `json:"return_date"`
}
