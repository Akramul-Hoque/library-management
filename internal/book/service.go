package book

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func AddBook(title, author string) (int, error) {
	b := Book{Title: title, Author: author}
	return save(b)
}

func GetAllBooks() ([]Book, error) {
	return findAll()
}
