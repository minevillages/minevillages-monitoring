package healthapi

import (
	"log"
	"os"

	"gopkg.in/mail.v2"
)

type Mail struct {
	Title string
	Body  string
}

func (m *Mail) Send() {
	mailAddress := os.Getenv("mail_address")
	mailPassword := os.Getenv("mail_password")
	mailReceipt := os.Getenv("mail_receipt")
	mailer := mail.NewMessage()
	mailer.SetHeader("From", mailAddress)
	mailer.SetHeader("To", mailReceipt)
	mailer.SetHeader("Subject", m.Title)
	mailer.SetBody("text/plain", m.Body)
	d := mail.NewDialer("smtp.gmail.com", 465, mailAddress, mailPassword)

	// Send the email
	if err := d.DialAndSend(mailer); err != nil {
		log.Fatalf("Failed to send monitor email: %v", err)
	}

	log.Println("Monitor Email sent successfully")
}
