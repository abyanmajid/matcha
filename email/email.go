package email

import (
	"bytes"
	"html/template"
	"net/smtp"
	"strings"
)

// Config holds SMTP server configuration details.
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
}

type Client struct {
	templateDir string
	config      Config
}

// NewClient initializes and returns a new email client.
func NewClient(config Config, templateDir string) *Client {
	return &Client{
		templateDir: templateDir,
		config:      config,
	}
}

// SendEmail sends an email with a custom sender.
func (c *Client) SendEmail(from string, to []string, subject string, templateName string, data interface{}) error {
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

	headers := map[string]string{
		"From":         from, // Allow passing custom sender
		"To":           strings.Join(to, ", "),
		"Subject":      subject,
		"MIME-Version": "1.0",
		"Content-Type": "text/html; charset=UTF-8",
	}

	var msg bytes.Buffer
	for key, value := range headers {
		msg.WriteString(key + ": " + value + "\r\n")
	}
	msg.WriteString("\r\n" + body.String())

	// Use authentication only if credentials are provided
	var auth smtp.Auth
	if c.config.Username != "" && c.config.Password != "" {
		auth = smtp.PlainAuth("", c.config.Username, c.config.Password, c.config.Host)
	}

	// Send email
	err = smtp.SendMail(c.config.Host+":"+c.config.Port, auth, from, to, msg.Bytes())
	if err != nil {
		return err
	}

	return nil
}
