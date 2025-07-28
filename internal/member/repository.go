package member

import (
	"library-management/pkg/db"
	"log"
)

func saveMember(m Member) error {
	query := `
		INSERT INTO members (name, contact, email, password, role, is_active)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := db.DB.Exec(query, m.Name, m.Contact, m.Email, m.Password, m.Role, m.IsActive)
	return err
}

func findAll() []Member {
	rows, err := db.DB.Query("SELECT id, name FROM members")
	if err != nil {
		log.Println("Error fetching members:", err)
		return nil
	}
	defer rows.Close()

	var members []Member
	for rows.Next() {
		var m Member
		rows.Scan(&m.ID, &m.Name)
		members = append(members, m)
	}
	return members
}
