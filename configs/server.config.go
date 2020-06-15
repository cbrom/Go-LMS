package configs

import (
	"time"
)

var CFG struct {
	Env    string `default:"dev" envconfig:"ENV"`
	Server struct {
		Host         string        `default:"0.0.0.0:3000" envconfig:"HOST"`
		Graphql      string        `default:"0.0.0.0:4000" envconfig:"GRAPHQL"`
		ReadTimeout  time.Duration `default:"10s" envconfig:"READ_TIMEOUT"`
		WriteTimeout time.Duration `default:"10s" envconfig:"WRITE_TIMOUT"`
	}

	App struct {
		Name        string `default:"Go-LMS-API" envconfig:"APP_Name"`
		EmailSender string `default:"test@go-lms.com" envconfig:"EMAIL_SENDER"`
	}

	Redis struct {
		Host            string        `default:"6379" envconfig:"REDIS_HOST"`
		DB              int           `default:"1" envconfig:"DB"`
		DialTimeout     time.Duration `default:"5s" envconfig:"DIAL_TIMOUT"`
		MaxmemoryPolicy string        `envconfig:"MAXMEMORY_POLICY"`
	}
	DB struct {
		Host       string `default:"127.0.0.1:5433" envconfig:"STORAGE_HOST"`
		User       string `default:"postgres" envconfig:"DB_HOST"`
		Pass       string `default:"postgres" envconfig:"DB_PASS" json:"-"`
		Database   string `default:"shared" envconfig:"DATABASE"`
		Driver     string `default:"postgres" envconfig:"Driver"`
		Timezone   string `default:"UTC" envconfig:"TIMEZONE"`
		DisableTLS bool   `default:"true" envconfig:"DISABLE_TLS"`
	}

	Auth struct {
		KeyExpiration time.Duration `deefault:"36000000s" envconfig:"KEY_EXPIRATION"`
	}
}
