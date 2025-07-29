package user

import (
	"library-management/pkg/db"
	"log"
)

func saveUser(u User) error {
	query := `
		INSERT INTO users (name, contact, email, password, role, is_active)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := db.DB.Exec(query, u.Name, u.Contact, u.Email, u.Password, u.Role, u.IsActive)
	return err
}

func findAllUsers() []User {
	rows, err := db.DB.Query(`SELECT id, name, contact, email, password, role, created_at, updated_at, is_active FROM users`)
	if err != nil {
		log.Println("Error fetching users:", err)
		return nil
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(
			&u.ID,
			&u.Name,
			&u.Contact,
			&u.Email,
			&u.Password,
			&u.Role,
			&u.CreatedAt,
			&u.UpdatedAt,
			&u.IsActive,
		)
		if err != nil {
			log.Println("Error scanning user row:", err)
			continue
		}
		users = append(users, u)
	}
	return users
}
