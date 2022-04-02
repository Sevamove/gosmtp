package email

import (
	"fmt"
	"log"
	"net/smtp"
)

//func plainAuth() smtp.Auth {
//	identity := auth.identity
//	username := auth.username
//	password := auth.password
//	hostname := auth.hostname
//
//	return smtp.PlainAuth(identity, username, password, hostname)
//}

func loginAuth() smtp.Auth {
	username := auth.username
	password := auth.password

	return &LoginAuth{username, password}
}

func SendEmail() {
	hostname := auth.hostname
	authentication := loginAuth()
	sender := sender.addr
	recipients := recipients.addr
	message := message.msg

	err := smtp.SendMail(
		hostname,
		authentication,
		sender,
		recipients,
		message,
	)

	if err != nil {
		log.Fatalln("Error!", err.Error())
	}

	fmt.Println("Awesome! Your email has been sent to the recipient.")
}
