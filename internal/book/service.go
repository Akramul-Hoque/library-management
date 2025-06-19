package book

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func AddBook(title, author string) {
	b := Book{Title: title, Author: author}
	save(b)
}

func GetAllBooks() []Book {
	return findAll()
}
