package main

import (
	"crypto/tls"
	"fmt"
	mail "gopkg.in/mail.v2"
)

func sendMail(resourceName string, changes string) {
	m := mail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "govardhan.pagidi@peerislands.com")

	// Set E-Mail receivers
	m.SetHeader("To", "govardhan.pagidi@mongodb.com")

	// Set E-Mail subject
	m.SetHeader("Subject", fmt.Sprintf("API Spec changes for %s", resourceName))

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", changes)

	// Settings for SMTP server
	d := mail.NewDialer("email-smtp.ap-northeast-1.amazonaws.com", 587, "AKIA2LKAQJBWFIOIRRWB", "BMjLdK8mgP7ORnCZXpNhbqtpsjaOkXN2fmD43SChEgy8")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return
}
