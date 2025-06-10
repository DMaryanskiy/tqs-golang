package tasks

import (
	"encoding/json"
	"log"

	"github.com/DMaryanskiy/tqs-golang/consumer/internal/processors"
)

func HandleTask(task Task) {
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
		processors.SendEmail(p.To, p.Subject, p.Body)
	default:
		log.Printf("Unknown task type: %s", task.Type)
	}
}
