package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) RegisterUser(u *User) error {
	// Validate required fields
	if u.Name == "" || u.Contact == "" || u.Password == "" {
		return errors.New("missing required fields")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	// Set default role if empty
	if u.Role == "" {
		u.Role = "USER"
	}
	u.IsActive = true

	return saveUser(*u)
}

func (s *Service) GetAllUsers() []User {
	return findAllUsers()
}

func (s *Service) GetUsersByType(userType string) []User {
	all := findAllUsers()
	var filtered []User
	for _, u := range all {
		if u.Role == userType {
			filtered = append(filtered, u)
		}
	}
	return filtered
}

func (s *Service) SearchUsers(contact, email, name, userType string) []User {
	all := findAllUsers()
	var filtered []User

	for _, u := range all {
		if contact != "" && u.Contact != contact {
			continue
		}
		if email != "" && u.Email != email {
			continue
		}
		if name != "" && u.Name != name {
			continue
		}
		if userType != "" && u.Role != userType {
			continue
		}
		filtered = append(filtered, u)
	}

	return filtered
}
