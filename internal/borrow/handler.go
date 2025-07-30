package borrow

import (
	"encoding/json"
	"library-management/internal/response"
	"net/http"
)

func BorrowBookHandler(w http.ResponseWriter, r *http.Request) {
	var req BorrowRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Universal(w, http.StatusBadRequest, false, "Invalid request payload", "INVALID_PAYLOAD", nil)
		return
	}
	err := BorrowBook(req.UserID, req.BookID)
	if err != nil {
		switch err.Error() {
		case "sql: no rows in result set":
			response.Universal(w, http.StatusConflict, false, "User already borrowed this book and has not returned it", "ALREADY_BORROWED", nil)
			return
		case "sql: transaction has already been committed or rolled back":
			response.Universal(w, http.StatusForbidden, false, "User cannot borrow more than 2 books at once", "BORROW_LIMIT", nil)
			return
		case "out of stock":
			response.Universal(w, http.StatusConflict, false, "No copies of this book are available", "OUT_OF_STOCK", nil)
			return
		}
		response.Universal(w, http.StatusInternalServerError, false, "Failed to borrow book", "BORROW_ERROR", nil)
		return
	}
	response.Universal(w, http.StatusOK, true, "Book borrowed successfully", "BORROW_SUCCESS", nil)
}

func ReturnBookHandler(w http.ResponseWriter, r *http.Request) {
	var req BorrowRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Universal(w, http.StatusBadRequest, false, "Invalid request payload", "INVALID_PAYLOAD", nil)
		return
	}
	err := ReturnBook(req.UserID, req.BookID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			response.Universal(w, http.StatusNotFound, false, "No active borrow found for this user and book", "NO_ACTIVE_BORROW", nil)
			return
		}
		response.Universal(w, http.StatusInternalServerError, false, "Failed to return book", "RETURN_ERROR", nil)
		return
	}
	response.Universal(w, http.StatusOK, true, "Book returned successfully", "RETURN_SUCCESS", nil)
}
