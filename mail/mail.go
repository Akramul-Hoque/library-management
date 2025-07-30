package mail

import (
	"log"
	"net/smtp"
	"os"
)

func SendOTPEmail(to, otp string) error {
	from := os.Getenv("MAIL_SENDER")
	password := os.Getenv("MAIL_APP_PASSWORD")

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	subject := "Subject: Your OTP Code\n"
	body := "Your OTP code is: " + otp
	message := []byte(subject + "\n" + body)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
	if err != nil {
		log.Println("Email sending error:", err)
		return err
	}

	log.Println("OTP email sent to", to)
	return nil
}
