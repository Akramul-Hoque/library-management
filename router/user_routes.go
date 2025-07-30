package router

import (
	"library-management/internal/middleware"
	"library-management/internal/user"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(r *mux.Router) {
	r.HandleFunc("/api/user/login", user.LoginHandler).Methods("POST")
	r.HandleFunc("/api/user", user.RegisterUserHandler).Methods("POST")
	r.HandleFunc("/api/user/otp-verify", user.OTPVerifyHandler).Methods("POST")
	// Protected routes:
	r.Handle("/api/users", middleware.AuthMiddleware(http.HandlerFunc(user.GetUsersHandler))).Methods("GET")
	r.Handle("/api/user/get-user", middleware.AuthMiddleware(http.HandlerFunc(user.SearchUsersHandler))).Methods("GET")
}
