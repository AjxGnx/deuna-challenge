package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/gommon/log"
)

type Config struct {
	ServerHost string `required:"true" split_words:"true" default:"localhost"`
	ServerPort int    `required:"true" split_words:"true" default:"8080"`
	StripeKey  string `required:"true" split_words:"true" default:""`
}

var once sync.Once
var config Config

func Environments() Config {
	once.Do(func() {
		if err := envconfig.Process("", &config); err != nil {
			log.Panicf("Error parsing environment vars %#v", err)
		}
	})

	return config
}
