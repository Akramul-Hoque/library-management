package main

import (
	"log"
	"net/http"

	"library-management/pkg/db"
	"library-management/router"
)

// @title Library Management API
// @version 1.0
// @description A REST API for managing books and members.
// @host localhost:8080
// @BasePath /api
func main() {
	db.Init()

	r := router.SetupRouter()

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
