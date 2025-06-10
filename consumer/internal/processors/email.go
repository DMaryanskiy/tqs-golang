package processors

import (
	"log"

	"github.com/DMaryanskiy/tqs-golang/consumer/internal/config"
	"gopkg.in/gomail.v2"
)

type Mailer interface {
	Send(msg *gomail.Message) error
}

type GomailDialer struct {
	Dialer *gomail.Dialer
}

func (g *GomailDialer) Send(msg *gomail.Message) error {
	return g.Dialer.DialAndSend(msg)
}

func SendEmail(mailer Mailer, to, subject, body string) error {
	cnf := config.Cfg
	message := gomail.NewMessage()

	message.SetHeader("From", cnf.EmailUser)
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)

	message.SetBody("text/plain", body)

	if err := mailer.Send(message); err != nil {
		log.Fatalf("Failed to send email: %v", err)
		return err
	} else {
		log.Println("Email was sent successfully")
		return nil
	}
}
