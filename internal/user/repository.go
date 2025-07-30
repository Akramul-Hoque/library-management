package user

import (
	"library-management/pkg/db"
	"log"
	"time"
)

func insertUser(u User) error {
	query := `
		INSERT INTO users (name, contact, email, password, role, is_active)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := db.DB.Exec(query, u.Name, u.Contact, u.Email, u.Password, u.Role, u.IsActive)
	return err
}

func updateUser(u User) error {
	query := `
		UPDATE users
		SET name = ?, contact = ?, email = ?, password = ?, role = ?, is_active = ?
		WHERE id = ?
	`
	_, err := db.DB.Exec(query, u.Name, u.Contact, u.Email, u.Password, u.Role, u.IsActive, u.ID)
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

func SaveOTP(userID int, otp string) error {
	query := "INSERT INTO user_otps (user_id, otp_code, expires_at) VALUES (?, ?, ?)"
	_, err := db.DB.Exec(query, userID, otp, time.Now().Add(5*time.Minute))
	if err != nil {
		log.Println("Error inserting book:", err)
		return err
	}
	return nil
}

func VerifyOTP(userID int, otp string) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM user_otps WHERE user_id = ? AND otp_code = ? AND expires_at > NOW() AND is_used = FALSE`
	err := db.DB.QueryRow(query, userID, otp).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func MarkOTPUsed(userID int, otp string) error {
	query := `UPDATE user_otps SET is_used = TRUE WHERE user_id = ? AND otp_code = ?`
	_, err := db.DB.Exec(query, userID, otp)
	return err
}
