package configs

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// AppConfig is the config structure
type AppConfig struct {
	TLS     TLS
	Storage Storage
	Address string `envconfig:"ADDRESS"`
}

//Storage is storage handler config
type Storage struct {
	HandlerName string `envconfig:"STORAGE_HANDLERNAME"`
	Host        string `envconfig:"STORAGE_HOST"`
	Port        string `envconfig:"STORAGE_PORT"`
	Dbuser      string `envconfig:"STORAGE_DBUSER"`
	Dbpassword  string `envconfig:"STORAGE_DBPASSWORD"`
	Database    string `envconfig:"STORAGE_DATABASE"`
	URL         string `envconfig:"STORAGE_URL"`
}

// TLS is the tls config for running the server
type TLS struct {
	Key   string `envconfig:"TLS_KEY"`
	Crt   string `envconfig:"TLS_CRT"`
	CACrt string `envconfig:"TLS_CACRT"`
}

//LoadConfig read env vars
func LoadConfig() (*AppConfig, error) {
	var conf AppConfig
	err := envconfig.Process("GO_LMS", &conf)
	if err != nil {
		return nil, err
	}
	if conf.Storage == (Storage{}) {
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
		tls := TLS{
			Key:   os.Getenv("TLS_KEY"),
			Crt:   os.Getenv("TLS_CRT"),
			CACrt: os.Getenv("TLS_CACRT"),
		}
		conf.Storage = storage
		conf.TLS = tls
		conf.Address = os.Getenv("ADDRESS")
	}

	return &conf, nil
}
