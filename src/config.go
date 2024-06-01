package src

import (
	"github.com/kelseyhightower/envconfig"
)

type (
	Config struct {
		ProjectName string `default:"contacts"`
		Port        string `envconfig:"PORT" default:"8085" required:"true"`
		Database    struct {
			URL string `envconfig:"DATABASE_URL" default:"postgres://svc_contact:contact_pwd@localhost:5432/contacts"`
		}
	}
)

func NewConfig() Config {
	var configs Config
	if err := envconfig.Process("", &configs); err != nil {
		panic(err.Error())
	}

	return configs
}
