package processors

import (
	"log"
	"strconv"

	"github.com/DMaryanskiy/tqs-golang/consumer/internal/config"
	"gopkg.in/gomail.v2"
)

func SendEmail(to, subject, body string) {
	cnf := config.Cfg
	message := gomail.NewMessage()

	message.SetHeader("From", cnf.EmailUser)
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)

	message.SetBody("text/plain", body)

	port, err := strconv.Atoi(cnf.EmailPort)
	if err != nil {
		log.Fatalf("Failed to convert port to int: %v", err)
	}

	dialer := gomail.NewDialer(cnf.EmailHost, port, cnf.EmailUser, cnf.EmailPass)
	if err := dialer.DialAndSend(message); err != nil {
		log.Fatalf("Failed to send email: %v", err)
	} else {
		log.Println("Email was sent successfully")
	}
}
