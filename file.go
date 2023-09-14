package main

import (
	"github.com/google/uuid"
	"os"
)

func writeUuid(fileName string) (*string, error) {
	str := uuid.New().String()
	if err := os.WriteFile(fileName, []byte(str), 0); err != nil {
		return nil, err
	}
	return &str, nil
}
