package router

import (
	"library-management/internal/book"
	"library-management/internal/middleware"
	"net/http"

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

	// @Summary List all books
	// @Description Retrieve books by name in the library
	// @Tags books
	// @Produce json
	// @Success 200 {array} book.Book
	// @Router /api/books [get]
	r.Handle("/api/books/search", middleware.AuthMiddleware(http.HandlerFunc(book.GetBooksByNameHandler))).Methods("GET")

	r.Handle("/api/books/edit", middleware.AuthMiddleware(http.HandlerFunc(book.EditBooksHandler))).Methods("PUT")
	r.Handle("/api/books/delete", middleware.AuthMiddleware(http.HandlerFunc(book.DeletetBooksHandler))).Methods("DELETE")
}
