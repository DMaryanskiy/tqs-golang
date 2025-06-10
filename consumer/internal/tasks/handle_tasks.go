package tasks

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/DMaryanskiy/tqs-golang/consumer/internal/config"
	"github.com/DMaryanskiy/tqs-golang/consumer/internal/processors"
	"gopkg.in/gomail.v2"
)

func HandleTask(task Task) {
	cnf := config.Cfg
	port, err := strconv.Atoi(cnf.EmailPort)
	if err != nil {
		log.Fatalf("Failed to convert port to int: %v", err)
	}
	mailer := &processors.GomailDialer{
		Dialer: gomail.NewDialer(
			cnf.EmailHost,
			port,
			cnf.EmailUser,
			cnf.EmailPass,
		),
	}

	switch task.Type {
	case TaskResizeImage:
		var p ResizeImagePayload
		err := json.Unmarshal(task.Payload, &p)
		if err != nil {
			log.Printf("Failed to load payload of image resizing: %s", err)
		}
		log.Println("Got image resize task")
		processors.ResizeImage(p.ImageURL, p.Width, p.Height)
	case TaskSendEmail:
		var p SendEmailPayload
		err := json.Unmarshal(task.Payload, &p)
		if err != nil {
			log.Printf("Failed to load payload of email sending: %s", err)
		}
		log.Println("Got send email task")
		processors.SendEmail(mailer, p.To, p.Subject, p.Body)
	default:
		log.Printf("Unknown task type: %s", task.Type)
	}
}
