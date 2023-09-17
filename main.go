package main

import (
	"golang.org/x/exp/slog"
	"os"
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

	git := NewGit(cfg)
	defer func() {
		out, _ := git.push()
		logger.Info(out)

	}()
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
		exit <- syscall.SIGHUP
	}()
	<-exit

}
