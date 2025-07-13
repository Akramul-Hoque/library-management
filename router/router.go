package router

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "library-management/docs"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Swagger docs
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	// Register routes
	RegisterBookRoutes(r)
	RegisterMemberRoutes(r)

	return r
}
