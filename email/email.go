package email

import (
	"bytes"
	"html/template"
	"net/smtp"
)

type SmtpConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}

type Client struct {
	templateDir string
	config      SmtpConfig
}

func NewClient(config SmtpConfig, templateDir string) *Client {
	return &Client{
		templateDir: templateDir,
		config:      config,
	}
}

func (c *Client) SendEmail(to []string, subject string, templateName string, data interface{}) error {
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
	headers["To"] = to[0]
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
