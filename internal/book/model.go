package book

// Book represents a book in the library
type Book struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	Published   string `json:"published"` // e.g. date string
	Publication string `json:"publication"`
	Quantity    int    `json:"quantity"`
}

// BookRequest is the payload for inserting a book
type BookRequest struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	Published   string `json:"published"`
	Publication string `json:"publication"`
	Quantity    int    `json:"quantity"`
}
