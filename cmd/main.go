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
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Book routes
	r.HandleFunc("/api/books", book.CreateBookHandler).Methods("POST")
	r.HandleFunc("/api/books", book.GetBooksHandler).Methods("GET")

	// Member routes
	r.HandleFunc("/api/members", member.RegisterMemberHandler).Methods("POST")
	r.HandleFunc("/api/members", member.GetMembersHandler).Methods("GET")

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
