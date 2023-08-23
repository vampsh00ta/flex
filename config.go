package main

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	GitUrl         string `json:"gitUrl" env:"gitUrl"`
	FileToReadName string `json:"fileToReadName" env:"fileToReadName"`
	FileDone       string `json:"fileDone" env:"fileDone"`
	SpotifyAlbum   string `json:"spotifyUrl" env:"spotifyAlbum"`
	Authorization  string `json:"authorization" env:"authorization"`
	User
}
type User struct {
	username string
	password string
}

func LoadCondig() Config {
	var cfg Config
	if err := cleanenv.ReadConfig(".env", &cfg); err != nil {
		panic(err)
	}
	return cfg
}
