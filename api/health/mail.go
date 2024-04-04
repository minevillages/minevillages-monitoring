package healthapi

import (
	"log"
	"os"

	"gopkg.in/mail.v2"
)

func SendMail(title string, body string) {
	mailAddress := os.Getenv("mail_address")
	mailPassword := os.Getenv("mail_password")
	mailReceipt := os.Getenv("mail_receipt")
	m := mail.NewMessage()
	m.SetHeader("From", mailAddress)
	m.SetHeader("To", mailReceipt)
	m.SetHeader("Subject", title)
	m.SetBody("text/plain", body)
	d := mail.NewDialer("smtp.gmail.com", 465, mailAddress, mailPassword)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		log.Fatalf("Failed to send monitor email: %v", err)
	}

	log.Println("Monitor Email sent successfully")
}
