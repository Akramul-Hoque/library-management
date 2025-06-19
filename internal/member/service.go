package member

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func RegisterMember(name string) {
	m := Member{Name: name}
	save(m)
}

func GetAllMembers() []Member {
	return findAll()
}
