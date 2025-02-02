package email

import (
	"bytes"
	"html/template"
	"net/smtp"
)

// Config holds SMTP server configuration details.
//
// Fields:
//
//	Host - SMTP server hostname
//	Port - SMTP server port
//	Username - SMTP username for authentication
//	Password - SMTP password for authentication
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
}

type client struct {
	templateDir string
	config      Config
}

// Newclient initializes and returns a new email client.
func NewClient(config Config, templateDir string) *client {
	return &client{
		templateDir: templateDir,
		config:      config,
	}
}

// SendEmail sends an email using a specified template and data.
func (c *client) SendEmail(to []string, subject string, templateName string, data interface{}) error {
	templatePath := c.templateDir + "/" + templateName + ".html"

	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	var body bytes.Buffer
	err = tmpl.Execute(&body, data)
	if err != nil {
		return err
	}

	headers := make(map[string]string)
	headers["From"] = c.config.Username
	headers["To"] = to[0] // only show the first recipient is shown in headers
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=UTF-8"

	var msg bytes.Buffer
	for key, value := range headers {
		msg.WriteString(key + ": " + value + "\r\n")
	}
	msg.WriteString("\r\n" + body.String())

	auth := smtp.PlainAuth("", c.config.Username, c.config.Password, c.config.Host)

	err = smtp.SendMail(c.config.Host+":"+c.config.Port, auth, c.config.Username, to, msg.Bytes())
	if err != nil {
		return err
	}

	return nil
}
