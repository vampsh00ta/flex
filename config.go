package main

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	GitUrl      string `json:"gitUrl" env:"gitUrl"`
	CommitCount int    `json:"commitCount" env:"commitCount"`
}

func LoadCondig() Config {
	var cfg Config
	if err := cleanenv.ReadConfig(".env", &cfg); err != nil {
		panic(err)
	}
	return cfg
}
