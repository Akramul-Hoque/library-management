package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	"library-management/internal/book"
	"library-management/internal/member"
	"library-management/pkg/db"

	_ "library-management/docs" // 👈 Needed for generated Swagger docs
)

// @title Library Management API
// @version 1.0
// @description A REST API for managing books and members.
// @host localhost:8080
// @BasePath /api
func main() {
	db.Init()

	r := mux.NewRouter()

	// Swagger route
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // Specify JSON endpoint
	))

	// Book routes
	// Book routes with Swagger annotations
	// @Summary Create a new book
	// @Description Add a new book to the library
	// @Tags books
	// @Accept json
	// @Produce json
	// @Param book body book.BookRequest true "Book details"
	// @Success 201 {object} book.Book
	// @Router /api/books [post]
	r.HandleFunc("/api/books", book.CreateBookHandler).Methods("POST")
	r.HandleFunc("/api/books", book.GetBooksHandler).Methods("GET")

	// Member routes
	r.HandleFunc("/api/members", member.RegisterMemberHandler).Methods("POST")
	r.HandleFunc("/api/members", member.GetMembersHandler).Methods("GET")

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
