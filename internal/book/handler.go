package book

import (
	"encoding/json"
	"library-management/response"
	"log"
	"net/http"
)

func CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	var req BookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, http.StatusBadRequest, "error", "Invalid JSON", "INVALID_PAYLOAD", nil)
		return
	}
	book := Book(req)
	err := AddBook(book)
	if err != nil {
		log.Println("DB insert error:", err)
		response.JSON(w, http.StatusInternalServerError, "error", "Failed to save book", "BOOK_SAVE_ERROR", nil)
		return
	}

	response.JSON(w, http.StatusCreated, "success", "Book created", "BOOK_CREATED", nil)
}

func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := GetAllBooks()
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "error", "Failed to fetch books", "BOOK_FETCH_ERROR", nil)
		return
	}
	response.JSON(w, http.StatusOK, "success", "Books retrieved", "BOOKS_RETRIEVED", books)
}

func GetBooksByNameHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		response.JSON(w, http.StatusBadRequest, "error", "Missing book name in query", "MISSING_QUERY", nil)
		return
	}

	books, err := GetBooksByName(name)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "error", "Failed to fetch books", "BOOK_FETCH_ERROR", nil)
		return
	}

	response.JSON(w, http.StatusOK, "success", "Books retrieved", "BOOKS_RETRIEVED", books)
}
