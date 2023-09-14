package main

import (
	"fmt"
	"log"
	"os/exec"
)

func gitinit(cfg *Config) (err error) {
	cmd := exec.Command("git", "init", ".")
	_, err = cmd.Output()
	if err != nil {
		return err
	}
	cmd = exec.Command("git", "remote", "add", "origin", cfg.GitUrl)
	_, err = cmd.Output()
	if err != nil {
		return err
	}
	cmd = exec.Command("git", "checkout", "-b", "main")
	_, err = cmd.Output()
	if err != nil {
		return err
	}
	log.Println("init")
	return err
}
func add(file *string) error {

	cmd := exec.Command("git", "add", *file)
	out, err := cmd.Output()

	if err != nil {
		return err
	}
	log.Println(string(out))
	return nil
}
func commit(text *string) error {
	cmd := exec.Command("git", "commit", "-m", *text)
	out, err := cmd.Output()
	if err != nil {
		return err
	}
	log.Println(string(out))
	return nil
}
func push() error {
	fmt.Println("pushed")
	return nil
}
