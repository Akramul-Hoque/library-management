package book

import (
	"library-management/pkg/db"
	"log"
)

func save(book Book) (int, error) {
	query := "INSERT INTO books (title, author) VALUES (?, ?)"
	result, err := db.DB.Exec(query, book.Title, book.Author)
	if err != nil {
		log.Println("Error inserting book:", err)
		return 0, err
	}

	id64, err := result.LastInsertId()
	if err != nil {
		log.Println("Error getting last insert id:", err)
		return 0, err
	}

	return int(id64), nil
}

func findAll() ([]Book, error) {
	query := "SELECT id, title, author FROM books"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author); err != nil {
			return nil, err
		}
		books = append(books, b)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}
