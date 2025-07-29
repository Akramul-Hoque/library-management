package router

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Swagger docs
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	// Register routes
	RegisterBookRoutes(r)
	RegisterUserRoutes(r)
	RegisterBorrowRoutes(r)

	return r
}
