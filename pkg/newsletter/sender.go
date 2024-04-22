package newsletter

import (
	"crypto/tls"

	"gopkg.in/mail.v2"
)

func SendNewsletterSender(emails []string, subject, body string) error {
	m := mail.NewMessage()

	m.SetHeader("From", "surathisin@yahoo.com")

	m.SetHeader("To", emails...)

	m.SetHeader("Subject", subject)

	m.SetBody("text/plain", body)

	d := mail.NewDialer("smtp.yourdomain.com", 587, "username", "password")

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
