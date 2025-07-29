package book

import (
	"library-management/pkg/db"
	"log"
)

func save(book Book) error {
	query := "INSERT INTO books (name, author, published, publication, quantity) VALUES (?, ?, ?, ?, ?)"
	_, err := db.DB.Exec(query, book.Name, book.Author, book.Published, book.Publication, book.Quantity)
	if err != nil {
		log.Println("Error inserting book:", err)
		return err
	}
	return nil
}

func findAll() ([]Book, error) {
	query := "SELECT name, author, published, publication, quantity FROM books"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.Name, &b.Author, &b.Published, &b.Publication, &b.Quantity); err != nil {
			return nil, err
		}
		books = append(books, b)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func findBooksByName(name string) ([]Book, error) {
	query := "SELECT name, author, published, publication, quantity FROM books WHERE name LIKE ?"
	rows, err := db.DB.Query(query, "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.Name, &b.Author, &b.Published, &b.Publication, &b.Quantity); err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}
