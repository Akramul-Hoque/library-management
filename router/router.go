package router

import (
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Register routes
	RegisterBookRoutes(r)
	RegisterUserRoutes(r)
	RegisterBorrowRoutes(r)

	return r
}
