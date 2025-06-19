package book

import (
	"library-management/pkg/db"
	"log"
)

func save(book Book) {
	query := "INSERT INTO books (title, author) VALUES (?, ?)"
	_, err := db.DB.Exec(query, book.Title, book.Author)
	if err != nil {
		log.Println("Error inserting book:", err)
	}
}

func findAll() []Book {
	rows, err := db.DB.Query("SELECT id, title, author FROM books")
	if err != nil {
		log.Println("Error fetching books:", err)
		return nil
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var b Book
		rows.Scan(&b.ID, &b.Title, &b.Author)
		books = append(books, b)
	}
	return books
}
