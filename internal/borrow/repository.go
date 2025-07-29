package borrow

import (
	"database/sql"
	"library-management/pkg/db"
	"time"
)

func BorrowBook(userID, bookID int) error {
	// Check if already borrowed and not returned
	var exists int
	err := db.DB.QueryRow(`SELECT COUNT(*) FROM borrowed_books WHERE user_id = ? AND book_id = ? AND is_returned = FALSE`, userID, bookID).Scan(&exists)
	if err != nil {
		return err
	}
	if exists > 0 {
		return sql.ErrNoRows // Use as duplicate borrow error
	}
	// Check if user has already borrowed 2 books and not returned
	var activeBorrows int
	err = db.DB.QueryRow(`SELECT COUNT(*) FROM borrowed_books WHERE user_id = ? AND is_returned = FALSE`, userID).Scan(&activeBorrows)
	if err != nil {
		return err
	}
	if activeBorrows >= 2 {
		return sql.ErrTxDone // Use as borrow limit error
	}
	// Check if book is available (quantity > 0)
	var quantity int
	err = db.DB.QueryRow(`SELECT quantity FROM books WHERE id = ?`, bookID).Scan(&quantity)
	if err != nil {
		return err
	}
	if quantity <= 0 {
		return sql.ErrNoRows // Use as out of stock error
	}
	// All checks passed, decrement quantity and insert borrow
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	_, err = tx.Exec(`UPDATE books SET quantity = quantity - 1 WHERE id = ? AND quantity > 0`, bookID)
	if err != nil {
		return err
	}
	_, err = tx.Exec(`INSERT INTO borrowed_books (user_id, book_id, borrow_date, is_returned) VALUES (?, ?, ?, FALSE)`, userID, bookID, time.Now())
	if err != nil {
		return err
	}
	err = tx.Commit()
	return err
}

func ReturnBook(userID, bookID int) error {
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	res, err := tx.Exec(`UPDATE borrowed_books SET is_returned = TRUE, return_date = ? WHERE user_id = ? AND book_id = ? AND is_returned = FALSE`, time.Now(), userID, bookID)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return sql.ErrNoRows // No active borrow found
	}
	// Increment book quantity
	_, err = tx.Exec(`UPDATE books SET quantity = quantity + 1 WHERE id = ?`, bookID)
	if err != nil {
		return err
	}
	err = tx.Commit()
	return err
}
