package user

import (
	"errors"
	"library-management/mail"
	"log"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) AuthenticateUser(contact, password string) (*User, error) {
	users := findAllUsers()
	for _, u := range users {
		if u.Contact == contact {
			if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
				break
			}

			if !u.IsActive {
				otp := generateOTP()
				err := SaveOTP(u.ID, otp)
				if err != nil {
					log.Printf("Failed to save OTP: %v", err)
					return nil, err
				}

				go func() {
					if sendErr := mail.SendOTPEmail(u.Email, otp); sendErr != nil {
						log.Printf("Failed to send OTP email: %v", sendErr)
					}
				}()

				return nil, ErrInactiveAccount{Message: "Account inactive. OTP sent for verification"}
			}
			return &u, nil
		}
	}
	return nil, errors.New("invalid credentials")
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

	return insertUser(*u)
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

func generateOTP() string {
	const digits = "0123456789"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	otp := make([]byte, 6)
	for i := range otp {
		otp[i] = digits[r.Intn(len(digits))]
	}
	return string(otp)
}

func (s *Service) VerifyOTP(contact, otp string) error {
	users := findAllUsers()
	var user *User
	for _, u := range users {
		if u.Contact == contact {
			user = &u
			break
		}
	}
	if user == nil {
		return errors.New("user not found")
	}

	valid, err := VerifyOTP(user.ID, otp)
	if err != nil {
		return err
	}
	if !valid {
		return errors.New("invalid or expired OTP")
	}

	err = MarkOTPUsed(user.ID, otp)
	if err != nil {
		return err
	}

	user.IsActive = true
	err = updateUser(*user)
	if err != nil {
		return err
	}

	return nil
}
