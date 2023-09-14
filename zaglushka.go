package main

import (
	"fmt"
	"github.com/google/uuid"
	"os/exec"
)

func createFile() (*string, error) {
	str := uuid.New().String()
	_, err := exec.Command("touch", str).Output()
	if err != nil {
		return nil, err
	}
	return &str, nil
}
func deleteFile(name *string) (err error) {
	file := fmt.Sprintf("%s", *name)
	_, err = exec.Command("rm", "-f", file).Output()
	if err != nil {
		return nil
	}
	return nil
}
