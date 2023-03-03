package common

import (
	"github.com/fancy-devops/fancy-devops-api/utils/setting"
	"gopkg.in/gomail.v2"
)

type Email struct {
}

func NewEmail() *Email {
	return &Email{}
}

func (obj *Email) SendEmail(mailTo, subject, body string) error {
	sec, err := setting.Cfg.GetSection("email")
	if err != nil {
		return err
	}

	host := sec.Key("HOST").String()
	port := sec.Key("PORT").MustInt()
	userName := sec.Key("USERNAME").String()
	senderName := sec.Key("SENDERNAME").String()
	authCode := sec.Key("AUTHCODE").String()

	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(userName, senderName))
	m.SetHeader("To", mailTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(host, port, userName, authCode)
	err = d.DialAndSend(m)
	return err
}
