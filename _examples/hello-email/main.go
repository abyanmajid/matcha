package main

import "github.com/abyanmajid/matcha/email"

func main() {
	emailClient := email.NewClient(email.Config{
		Host:     "localhost",
		Port:     "1025",
		Username: "mailhog",
		Password: "mailhog",
	}, "_examples/hello-email/templates")

	type EmailData struct {
		Name string
	}

	emailClient.SendEmail([]string{"abyan@abydyl.net"}, "Hello World!", "hello", EmailData{
		Name: "Abyan Majid",
	})
}
