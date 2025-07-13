package book

import (
	"encoding/json"
	"library-management/response"
	"log"
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
		response.JSON(w, http.StatusBadRequest, "error", "Invalid JSON", "INVALID_PAYLOAD", nil)
		return
	}

	id, err := AddBook(b)
	if err != nil {
		log.Println("DB insert error:", err)
		response.JSON(w, http.StatusInternalServerError, "error", "Failed to save book", "BOOK_SAVE_ERROR", nil)
		return
	}

	b.ID = id
	response.JSON(w, http.StatusCreated, "success", "Book created", "BOOK_CREATED", nil) // Send back created book
}

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
func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := GetAllBooks()
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "error", "Failed to fetch books", "BOOK_FETCH_ERROR", nil)
		return
	}
	response.JSON(w, http.StatusOK, "success", "Books retrieved", "BOOKS_RETRIEVED", books)
}

func GetBooksByNameHandler(w http.ResponseWriter, r *http.Request) {
	// Get 'name' from query string
	name := r.URL.Query().Get("name")
	if name == "" {
		response.JSON(w, http.StatusBadRequest, "error", "Missing book name in query", "MISSING_QUERY", nil)
		return
	}

	// Call a service/repo function to search by name
	books, err := GetBooksByName(name)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "error", "Failed to fetch books", "BOOK_FETCH_ERROR", nil)
		return
	}

	response.JSON(w, http.StatusOK, "success", "Books retrieved", "BOOKS_RETRIEVED", books)
}
