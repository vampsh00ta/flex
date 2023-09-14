package main

import (
	"github.com/xlab/closer"
	"golang.org/x/exp/slog"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func main() {
	cfg := LoadCondig()
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	closer.Bind(func() {
		cmd := exec.Command("git", "push", "origin", "main")
		result, _ := cmd.Output()
		logger.Info(string(result))

	})

	gitinit(&cfg)
	go func() {

		for i := 0; i < cfg.CommitCount; i++ {

			file, err := createFile()
			if err != nil {
				logger.Info(err.Error())

			}

			if err := add(); err != nil {
				logger.Info(err.Error())

			}
			if err := commit(file); err != nil {
				logger.Info(err.Error())

			}
			if err := deleteFile(file); err != nil {
				logger.Info(err.Error())

			}

		}
		os.Exit(0)
	}()
	<-exit
	closer.Hold()

}
