package main

import (
	"fmt"
	"github.com/xlab/closer"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func main() {
	cfg := LoadCondig()
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	closer.Bind(func() {

		cmd := exec.Command("git", "push", "origin", "main")
		cmd.Output()
		fmt.Println("slatt")

	})

	gitinit(&cfg)
	//if err := ; err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	go func() {

		for {

			file, err := createFile()
			if err != nil {
			}

			if err := add(); err != nil {

			}
			if err := commit(file); err != nil {

			}
			if err := deleteFile(file); err != nil {

			}

		}
	}()
	<-exit
	closer.Hold()

}
