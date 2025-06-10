package tasks

import "encoding/json"

type TaskType string

const (
	TaskResizeImage TaskType = "resize_image"
	TaskSendEmail   TaskType = "send_email"
)

type Task struct {
	Type    TaskType        `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type ResizeImagePayload struct {
	ImageURL string `json:"image_url"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
}

type SendEmailPayload struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}
