package member

import (
	"library-management/pkg/db"
	"log"
)

func save(member Member) {
	query := "INSERT INTO members (name) VALUES (?)"
	_, err := db.DB.Exec(query, member.Name)
	if err != nil {
		log.Println("Error inserting member:", err)
	}
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
