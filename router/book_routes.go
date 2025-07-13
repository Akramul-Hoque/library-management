package router

import (
	"library-management/internal/book"

	"github.com/gorilla/mux"
)

func RegisterBookRoutes(r *mux.Router) {
	// @Summary Create a new book
	// @Description Add a new book to the library
	// @Tags books
	// @Accept json
	// @Produce json
	// @Param book body book.BookRequest true "Book details"
	// @Success 201 {object} book.Book
	// @Router /api/books [post]
	r.HandleFunc("/api/books", book.CreateBookHandler).Methods("POST")

	// @Summary List all books
	// @Description Retrieve all books in the library
	// @Tags books
	// @Produce json
	// @Success 200 {array} book.Book
	// @Router /api/books [get]
	r.HandleFunc("/api/books", book.GetBooksHandler).Methods("GET")
}
