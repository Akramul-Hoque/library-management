package member

import "time"

type Member struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Contact   string    `json:"contact"` // <- must exist
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	Role      string    `json:"role,omitempty"` // <- must exist
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	IsActive  bool      `json:"is_active,omitempty"`
}
