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
