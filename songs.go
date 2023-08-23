package main

import (
	"io"
	"os"
	"strings"
)

const lineBreak = "\n"

func getSongs(count int, fileName string) ([]string, error) {
	songsBytes, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	b, err := io.ReadAll(songsBytes)
	songs := strings.Split(string(b), lineBreak)
	if err = songsBytes.Close(); err != nil {
		return nil, err
	}
	return songs, err

}
