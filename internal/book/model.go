package book

// Book represents a book in the library
type Book struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Published   string `json:"published"`
	Publication string `json:"publication"`
	Quantity    int    `json:"quantity"`
}

type BookRequest struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	Published   string `json:"published"`
	Publication string `json:"publication"`
	Quantity    int    `json:"quantity"`
}

type BookEditRequest struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Published   string `json:"published"`
	Publication string `json:"publication"`
	Quantity    int    `json:"quantity"`
}

type BookDeleteRequest struct {
	Id int `json:"id"`
}
