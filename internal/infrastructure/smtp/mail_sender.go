package smtp

import (
	"bytes"
	"fmt"
	"net/smtp"

	"github.com/Team-Work-Forever/FireWatchRest/config"
)

type Mail struct {
	To      string
	Subject string
	Body    string
}

func New(to string, subject string, body string) *Mail {
	return &Mail{
		To:      to,
		Subject: subject,
		Body:    body,
	}
}

func (m *Mail) Send() error {
	env := config.GetCofig()
	auth := smtp.PlainAuth("", env.SMTP_HOST_USER, env.SMTP_HOST_PASSWORD, env.SMTP_HOST)
	to := []string{m.To}

	var message bytes.Buffer
	message.WriteString("From: " + env.SMTP_HOST_EMAIL + "\r\n")
	message.WriteString("To: " + to[0] + "\r\n")
	message.WriteString("Subject: " + m.Subject + "\r\n")
	message.WriteString("MIME-version: 1.0;\r\n")
	message.WriteString("Content-Type: text/html; charset=\"UTF-8\";\r\n\r\n")
	message.WriteString(m.Body)

	addr := fmt.Sprintf("%s:%s", env.SMTP_HOST, env.SMTP_PORT)
	return smtp.SendMail(addr, auth, env.SMTP_HOST_EMAIL, to, message.Bytes())
}
