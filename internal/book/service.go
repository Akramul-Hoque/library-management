package book

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func AddBook(b Book) error {
	return save(b)
}

func GetAllBooks() ([]Book, error) {
	return findAll()
}

func GetBooksByName(name string) ([]Book, error) {
	return findBooksByName(name)
}
