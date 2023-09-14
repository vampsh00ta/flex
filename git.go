package main

import (
	"os/exec"
)

type Git struct {
	cfg Config
}

func NewGit(cfg Config) *Git {
	cmd := exec.Command("git", "init", ".")
	cmd.Output()

	cmd = exec.Command("git", "remote", "add", "origin", cfg.GitUrl)
	cmd.Output()

	cmd = exec.Command("git", "remote", "set-url", "origin", cfg.GitUrl)
	cmd.Output()

	cmd = exec.Command("git", "checkout", "-b", "main")
	cmd.Output()

	return &Git{cfg}
}

func (git Git) add() (string, error) {

	cmd := exec.Command("git", "add", git.cfg.TextFile)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
func (git Git) commit(text *string) (string, error) {
	cmd := exec.Command("git", "commit", "-m", *text)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil

}
func (git Git) push() (string, error) {
	cmd := exec.Command("git", "push", "origin", "main")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
