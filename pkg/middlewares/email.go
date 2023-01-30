package middlewares

import (
	"bytes"
	"crypto/tls"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"go-lms-of-pupilfirst/cmd/models"
	"go-lms-of-pupilfirst/configs"

	"github.com/k3a/html2text"
	"gopkg.in/gomail.v2"
)

type EmailData struct {
	URL       string
	FirstName string
	Subject   string
}

// ? Email template parser

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}

func SendEmail(user *models.User, data *EmailData) {
	config, err := configs.LoadConfig()

	if err != nil {
		log.Fatal("could not load config", err)
	}

	// Sender data.
	from := config.Email.EmailFrom
	smtpPass := config.Email.SMTPPass
	smtpUser := config.Email.SMTPUser
	to := user.Email
	smtpHost := config.Email.SMTPHost
	smtpPort := config.Email.SMTPPort
	a, err := strconv.Atoi(smtpPort)
	if err != nil {
		log.Fatal("Could not convert to int ", err)
	}

	var body bytes.Buffer

	template, err := ParseTemplateDir("user-interface/templates")
	if err != nil {
		log.Fatal("Could not parse template", err)
	}

	template.ExecuteTemplate(&body, "verificationCode.html", &data)

	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(smtpHost, a, smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send Email
	if err := d.DialAndSend(m); err != nil {
		log.Fatal("Could not send email: ", err)
	}

}
