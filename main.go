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
	git := NewGit(cfg)

	go func() {
		var out string
		var uuid *string

		var err error
		for i := 0; i < cfg.CommitCount; i++ {
			uuid, err = writeUuid(cfg.TextFile)
			if err != nil {
				logger.Info(err.Error())
			}

			out, err = git.add()
			if err != nil {
				logger.Info(err.Error())
			}
			logger.Info(out)
			out, err = git.commit(uuid)
			if err != nil {
				logger.Info(err.Error())
			}
			logger.Info(out)

		}
		os.Exit(0)
	}()
	closer.Hold()

}
