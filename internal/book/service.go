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

func EditBook(b Book) error {
	return updateBook(b)
}
func DeleteBook(id int) error {
	return deleteBook(id)
}
