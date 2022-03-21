package config

import (
	"fmt"

	"github.com/Netflix/go-env"
)

type Environment struct {
	ApiPort string `env:"API_PORT"`

	Database struct {
		Host     string `env:"HOST"`
		User     string `env:"USER"`
		Password string `env:"PASSWORD"`
		DbName   string `env:"DB_NAME"`
		Port     string `env:"PORT"`
		SslMode  string `env:"SSL_MODE"`
		TimeZone string `env:"TIME_ZONE"`
	}
}

var Env Environment

func LoadEnv() {
	_, err := env.UnmarshalFromEnviron(&Env)
	if err != nil {
		fmt.Print(err.Error())
	}
}
