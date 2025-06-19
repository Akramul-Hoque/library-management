package book

import (
	"encoding/json"
	"net/http"
)

// CreateBookHandler godoc
// @Summary Add a new book
// @Description Create a new book by providing title and author
// @Tags books
// @Accept json
// @Produce json
// @Param book body Book true "Book JSON"
// @Success 201 {object} Book
// @Failure 400 {string} string "Bad Request"
// @Router /api/books [post]
func CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	var b Book
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	AddBook(b.Title, b.Author)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(b)
}

// GetBooksHandler godoc
// @Summary Get all books
// @Description Retrieve a list of all books
// @Tags books
// @Produce json
// @Success 200 {array} Book
// @Router /api/books [get]
func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	books := GetAllBooks()
	json.NewEncoder(w).Encode(books)
}
