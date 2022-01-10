package helper

import (
	"fmt"
	"net/smtp"
)

func Mail() {

	// Sender data.
	from := "omonadefranklyn@gmail.com"
	password := "rutheze4me"

	// Receiver email address.
	to := []string{
		"franklyn.omonade@decagon.dev",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "465"

	// Message.
	message := []byte("This is a test email message.")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
