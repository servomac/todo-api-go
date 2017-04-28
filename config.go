package main

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port     int64  `default:"8000"`
	Database string `default:"todo.db"`
}

func getConfig() (c Config) {
	err := envconfig.Process("app", &c)
	if err != nil {
		log.Fatal("Env", err.Error())
	}
	return
}
