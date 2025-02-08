package main

import "github.com/abyanmajid/matcha/email"

func main() {
	emailClient := email.NewClient(email.Config{
		Host:     "localhost",
		Port:     "1025",
		Username: "mailhog",
		Password: "mailhog",
	}, "_examples/email-hello/templates")

	type EmailData struct {
		Name string
	}

	emailClient.SendEmail("sender@example.com", []string{"abyan@abydyl.net"}, "Hello World!", "hello", EmailData{
		Name: "Abyan Majid",
	})
}
