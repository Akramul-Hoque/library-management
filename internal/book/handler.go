package book

import (
	"encoding/json"
	"library-management/internal/response"
	"log"
	"net/http"
)

func CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	var req BookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, http.StatusBadRequest, "error", "Invalid JSON", "INVALID_PAYLOAD", nil)
		return
	}
	book := Book{
		Name:        req.Name,
		Author:      req.Author,
		Published:   req.Published,
		Publication: req.Publication,
		Quantity:    req.Quantity,
	}
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

func EditBooksHandler(w http.ResponseWriter, r *http.Request) {
	var req BookEditRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, http.StatusBadRequest, "error", "Invalid JSON", "INVALID_PAYLOAD", nil)
		return
	}

	book := Book(req)
	if err := EditBook(book); err != nil {
		switch err.Error() {
		case "sql: no rows in result set":
			response.Universal(w, http.StatusConflict, false, "Books not fOund", "BOOK_FETCH_ERROR", nil)
			return
		}
		response.JSON(w, http.StatusOK, "success", "Book updated", "BOOK_UPDATED", nil)
	}
}
func DeletetBooksHandler(w http.ResponseWriter, r *http.Request) {
	var req BookDeleteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, http.StatusBadRequest, "error", "Invalid JSON", "INVALID_PAYLOAD", nil)
		return
	}

	if err := DeleteBook(req.Id); err != nil {
		switch err.Error() {
		case "sql: no rows in result set":
			response.Universal(w, http.StatusConflict, false, "Books not found", "BOOK_FETCH_ERROR", nil)
			return
		}
		response.JSON(w, http.StatusInternalServerError, "error", "Failed to delete book", "BOOK_DELETE_ERROR", nil)
		return
	}

	response.JSON(w, http.StatusOK, "success", "Book deleted successfully", "BOOK_DELETED", nil)
}
