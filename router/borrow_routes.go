package router

import (
	"library-management/internal/borrow"
	"library-management/internal/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterBorrowRoutes(r *mux.Router) {
	r.Handle("/api/borrow", middleware.AuthMiddleware(http.HandlerFunc(borrow.BorrowBookHandler))).Methods("POST")
	r.Handle("/api/return", middleware.AuthMiddleware(http.HandlerFunc(borrow.ReturnBookHandler))).Methods("POST")
}
