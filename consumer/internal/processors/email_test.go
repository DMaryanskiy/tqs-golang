package processors_test

import (
	"os"
	"testing"

	"github.com/DMaryanskiy/tqs-golang/consumer/internal/config"
	"github.com/DMaryanskiy/tqs-golang/consumer/internal/processors"
	"gopkg.in/gomail.v2"
)

// Mock implementation
type MockMailer struct {
	Called bool
}

func (m *MockMailer) Send(msg *gomail.Message) error {
	m.Called = true
	return nil // simulate success
}

func init() {
	_ = os.Setenv("EMAIL_HOST", "smtp.gmail.com")
	_ = os.Setenv("EMAIL_PORT", "587")
	_ = os.Setenv("EMAIL_USER", "your_email@example.com")
	_ = os.Setenv("EMAIL_PASS", "your_smtp_password")
	config.NewConfig()
}

func TestSendEmailWithMock(t *testing.T) {
	mock := &MockMailer{}
	err := processors.SendEmail(mock, "recipient@example.com", "Test Subject", "Test Body")

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if !mock.Called {
		t.Errorf("Expected Send to be called on mock mailer")
	}
}
