package handlers

import (
	"client/model/web"
	"client/service"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	bookService *service.BookService
}

func NewBookHandler() (*BookHandler, error) {
	bookService, err := service.NewBookService()
	if err != nil {
		return nil, fmt.Errorf("could not initialize BookService: %w", err)
	}
	return &BookHandler{bookService: bookService}, nil
}

// RegisterBook handles the book registration HTTP request
// @Summary Register a new book
// @Description Register a new book with title, author, and published date
// @Tags books
// @Accept json
// @Produce json
// @Param request body web.RegisterBookRequest true "Register Book Request"
// @Success 200 {object} pb.BookResponse "Book registered successfully"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security BearerAuth
// @Router /books/register [post]
func (h *BookHandler) RegisterBook(ctx echo.Context) error {
	var req web.RegisterBookRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	jwtToken := ctx.Request().Header.Get("Authorization")
	if jwtToken == "" {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{"error": "JWT token is missing"})
	}

	resp, err := h.bookService.RegisterBook(req.Title, req.Author, req.PublishedDate, jwtToken)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, resp)
}

// BorrowBook handles the borrow book HTTP request
// @Summary Borrow a book
// @Description Borrow a book by providing book ID, user ID, and borrowed date
// @Tags books
// @Accept json
// @Produce json
// @Param request body web.BorrowBookRequest true "Borrow Book Request"
// @Success 200 {object} pb.BorrowedBookResponse "Book borrowed successfully"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security BearerAuth
// @Router /books/borrow [post]
func (h *BookHandler) BorrowBook(ctx echo.Context) error {
	var req web.BorrowBookRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	jwtToken := ctx.Request().Header.Get("Authorization")
	if jwtToken == "" {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{"error": "JWT token is missing"})
	}

	resp, err := h.bookService.BorrowBook(req.BookID, req.UserID, req.BorrowedDate, jwtToken)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, resp)
}

// ReturnBook handles the return book HTTP request
// @Summary Return a book
// @Description Return a book by providing book ID, user ID, and return date
// @Tags books
// @Accept json
// @Produce json
// @Param request body web.ReturnBookRequest true "Return Book Request"
// @Success 200 {object} pb.BorrowedBookResponse "Book returned successfully"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security BearerAuth
// @Router /books/return [post]
func (h *BookHandler) ReturnBook(ctx echo.Context) error {
	var req web.ReturnBookRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	jwtToken := ctx.Request().Header.Get("Authorization")
	if jwtToken == "" {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{"error": "JWT token is missing"})
	}

	resp, err := h.bookService.ReturnBook(req.BookID, req.UserID, req.ReturnDate, jwtToken)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, resp)
}