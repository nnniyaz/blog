package email

import (
	"gopkg.in/gomail.v2"
)

type Email interface {
	SendMail(to []string, subject string, htmlBody string) error
}

type email struct {
	dialer *gomail.Dialer
}

func New(host string, port int, user string, pass string) (Email, error) {
	d := gomail.NewDialer(host, port, user, pass)
	if err := d.DialAndSend(); err != nil {
		return nil, err
	}
	return &email{dialer: d}, nil
}

func (e *email) SendMail(to []string, subject string, htmlBody string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.dialer.Username)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", htmlBody)
	return e.dialer.DialAndSend(m)
}
