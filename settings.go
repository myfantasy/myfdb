package main

import (
	"os"
	"time"

	env "github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

var (
	config envConfig
)

type envConfig struct {
	DBFolder       string        `default:"db/"`
	DBFlushTimeout time.Duration `default:"1s"`

	Port int `default:"9170"`

	LogLevel log.Level `envconfig:"LOG_LEVEL" default:"info"`
}

func initConf() {
	if err := env.Process("", &config); err != nil {
		env.Usage("", &config)
		os.Exit(1)
	}

	log.SetLevel(config.LogLevel)

}
