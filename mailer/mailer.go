package mailer

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

// Mailer 提供邮件发送功能
type Mailer struct {
	SMTPHost     string
	SMTPPort     int
	SMTPUser     string
	SMTPPassword string
}

// NewMailer 创建新的 Mailer 实例
func NewMailer(host string, port int, user, password string) *Mailer {
	return &Mailer{
		SMTPHost:     host,
		SMTPPort:     port,
		SMTPUser:     user,
		SMTPPassword: password,
	}
}

// SendEmail 使用动态邮箱发送邮件
func (m *Mailer) SendEmail(recipient, subject, body string) error {
	message := gomail.NewMessage()
	message.SetHeader("From", m.SMTPUser)
	message.SetHeader("To", recipient)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", body)

	dialer := gomail.NewDialer(m.SMTPHost, m.SMTPPort, m.SMTPUser, m.SMTPPassword)
	if err := dialer.DialAndSend(message); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	return nil
}
