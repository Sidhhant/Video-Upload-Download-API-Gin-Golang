package config

import (
	"github.com/pkg/errors"

	"github.com/kelseyhightower/envconfig"
)

const (
	envDevelopment = "development"
	envProduction  = "production"
)

type Env struct {

	// Env is environment where application is running.The value must be
	// "development" or "production".
	Env string `envconfig:"ENV" default:"development"`

	HTTPPort string `envconfig:"HTTP_PORT" default:"8080"`

	ContextDB string `envconfig:"DB_NAME" default:"dbname"`

	DBPassword string `envconfig:"DB_PASSWORD" default:"dbpassword"`

	DBUser string `envconfig:"DB_USER" default:"dbuser"`
}

// IsProduction returns true if it is production environment
func (e *Env) IsProduction() bool {
	return e.Env == envProduction
}

// ReadFromEnv reads configuration from environmental variables
// defined by Env struct.
func ReadFromEnv() (*Env, error) {
	var env Env
	if err := envconfig.Process("", &env); err != nil {
		return nil, errors.Wrap(err, "failed to process envconfig")
	}

	return &env, nil
}
