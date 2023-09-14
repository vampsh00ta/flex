package main

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	GitUrl string `json:"gitUrl" env:"gitUrl"`
}

func LoadCondig() Config {
	var cfg Config
	if err := cleanenv.ReadConfig(".env", &cfg); err != nil {
		panic(err)
	}
	return cfg
}
