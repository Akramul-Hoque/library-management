package member

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) RegisterMember(m *Member) error {
	// Validate required fields
	if m.Name == "" || m.Contact == "" || m.Password == "" {
		return errors.New("missing required fields")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	m.Password = string(hashedPassword)

	// Set default role if empty
	if m.Role == "" {
		m.Role = "USER"
	}
	m.IsActive = true

	return saveMember(*m) // assuming saveMember accepts a value, adjust if pointer
}

func (s *Service) GetAllMembers() []Member {
	return findAll()
}
