package newsletter

import (
	"crypto/tls"

	"gopkg.in/mail.v2"
)

func SendNewsletterSender(emails []string, subject, body string) error {
	m := mail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "surathisin@yahoo.com")

	// Set E-Mail receivers
	m.SetHeader("To", emails...)

	// Set E-Mail subject
	m.SetHeader("Subject", subject)

	// Set E-Mail body
	m.SetBody("text/plain", body)

	// SMTP server configuration
	d := mail.NewDialer("smtp.yourdomain.com", 587, "username", "password")

	// Secure configuration for SMTP
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
