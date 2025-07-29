-- +goose Up
CREATE TABLE IF NOT EXISTS books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    published VARCHAR(50) NOT NULL,
    publication VARCHAR(255) NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS books;

