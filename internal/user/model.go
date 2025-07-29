package user

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Contact   string    `json:"contact"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	Role      string    `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	IsActive  bool      `json:"is_active,omitempty"`
}

type UserResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Contact   string `json:"contact"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	IsActive  bool   `json:"is_active"`
}

type loginRequest struct {
	Contact  string `json:"contact"`
	Password string `json:"password"`
}
