package book

// Book represents a book in the library
// swagger:model
type Book struct {
	ID     int    `json:"id" example:"1"`
	Title  string `json:"title" example:"The Go Programming Language"`
	Author string `json:"author" example:"Alan A. Donovan"`
}

// swagger:model
type BookRequest struct {
	Title  string `json:"title" example:"Clean Code"`
	Author string `json:"author" example:"Robert C. Martin"`
}
