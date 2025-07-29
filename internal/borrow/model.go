package borrow

type BorrowedBook struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	BookID     int    `json:"book_id"`
	BorrowDate string `json:"borrow_date"`
	ReturnDate string `json:"return_date,omitempty"`
	IsReturned bool   `json:"is_returned"`
}

type BorrowRequest struct {
	UserID int `json:"user_id"`
	BookID int `json:"book_id"`
}
