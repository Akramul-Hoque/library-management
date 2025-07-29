-- +goose Up
CREATE TABLE borrowed_books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    book_id INT NOT NULL,
    borrow_date DATETIME NOT NULL,
    return_date DATETIME,
    is_returned BOOLEAN NOT NULL DEFAULT FALSE,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_book FOREIGN KEY (book_id) REFERENCES books(id),
    UNIQUE KEY unique_borrow (user_id, book_id, is_returned)
);
-- +goose Down
DROP TABLE borrowed_books;
