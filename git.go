package main

import (
	"fmt"
	"log"
	"os/exec"
)

func gitinit(cfg Config) (err error) {
	_, err = exec.Command("git init .").Output()
	if err != nil {
		return nil
	}
	_, err = exec.Command("git remote add origin %s", cfg.GitUrl).Output()
	if err != nil {
		return nil
	}
	log.Println("init")
	return err
}
func add() error {

	cmd := exec.Command("git", "add", ".")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("could not run command: ", err)
	}

	if err != nil {
		return err
	}
	log.Println(string(out))
	return nil
}
func commit(text string) error {
	cmd := exec.Command("git", "commit", "-m", text)
	out, err := cmd.Output()
	if err != nil {
		return err
	}
	log.Println(string(out))
	return nil
}
