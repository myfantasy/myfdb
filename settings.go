package main

import (
	"os"
	"time"

	env "github.com/kelseyhightower/envconfig"
)

var (
	config envConfig
)

type envConfig struct {
	DBFolder       string        `default:"db/"`
	DBFlushTimeout time.Duration `default:"1s"`

	Port int `default:"9170"`
}

func initConf() {
	if err := env.Process("", &config); err != nil {
		env.Usage("", &config)
		os.Exit(1)
	}

}
