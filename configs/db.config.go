package configs

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// AppConfig is the config structure
type AppConfig struct {
	Email        Email
	TLS          TLS
	Storage      Storage
	Address      string `envconfig:"ADDRESS"`
	ClientOrigin string `envconfig:"CLIENT_ORIGIN"`
}

// Storage is storage handler config
type Storage struct {
	HandlerName string `envconfig:"STORAGE_HANDLERNAME"`
	Host        string `envconfig:"STORAGE_HOST"`
	Port        string `envconfig:"STORAGE_PORT"`
	Dbuser      string `envconfig:"STORAGE_DBUSER"`
	Dbpassword  string `envconfig:"STORAGE_DBPASSWORD"`
	Database    string `envconfig:"STORAGE_DATABASE"`
	URL         string `envconfig:"STORAGE_URL"`
}

type Email struct {
	EmailFrom string `envconfig:"EMAIL_FROM"`
	SMTPHost  string `envconfig:"SMTP_HOST"`
	SMTPPort  string `envconfig:"SMTP_PORT"`
	SMTPUser  string `envconfig:"SMTP_USER"`
	SMTPPass  string `envconfig:"SMTP_PASS"`
}

// TLS is the tls config for running the server
type TLS struct {
	Key   string `envconfig:"TLS_KEY"`
	Crt   string `envconfig:"TLS_CRT"`
	CACrt string `envconfig:"TLS_CACRT"`
}

// LoadConfig read env vars
func LoadConfig() (*AppConfig, error) {
	var conf AppConfig
	err := envconfig.Process("GO_LMS", &conf)
	if err != nil {
		return nil, err
	}
	if conf.Storage.Host == "db" || conf.Storage == (Storage{}) {
		err = godotenv.Load()
		if err != nil {
			return nil, err
		}

		storage := Storage{
			HandlerName: os.Getenv("STORAGE_HANDLERNAME"),
			Host:        os.Getenv("STORAGE_HOST"),
			URL:         os.Getenv("STORAGE_URL"),
			Database:    os.Getenv("STORAGE_DATABASE"),
			Port:        os.Getenv("STORAGE_PORT"),
			Dbuser:      os.Getenv("STORAGE_DBUSER"),
			Dbpassword:  os.Getenv("STORAGE_DBPASSWORD"),
		}
		// If running in docker-compose, we should check if STORAGE_HOST was set.
		if conf.Storage.Host == "db" {
			storage.Host = conf.Storage.Host
		}
		email := Email{
			EmailFrom: os.Getenv("EMAIL_FROM"),
			SMTPHost:  os.Getenv("SMTP_HOST"),
			SMTPPass:  os.Getenv("SMTP_PASS"),
			SMTPPort:  os.Getenv("SMTP_PORT"),
			SMTPUser:  os.Getenv("SMTP_USER"),
		}

		tls := TLS{
			Key:   os.Getenv("TLS_KEY"),
			Crt:   os.Getenv("TLS_CRT"),
			CACrt: os.Getenv("TLS_CACRT"),
		}
		conf.Email = email
		conf.Storage = storage
		conf.TLS = tls
		conf.Address = os.Getenv("ADDRESS")
		conf.ClientOrigin = os.Getenv("CLIENT_ORIGIN")
	}

	return &conf, nil
}
