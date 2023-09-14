package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"time"
)

func main() {
	cfg := LoadCondig()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)

	gitinit(&cfg)
	go func() {

		for {

			file, err := createFile()
			if err != nil {
				os.Exit(1)
			}

			if err := add(); err != nil {

				os.Exit(1)
			}
			if err := commit(file); err != nil {

				os.Exit(1)
			}
			if err := deleteFile(file); err != nil {
				os.Exit(1)

			}

		}
	}()
	<-exit
	time.Sleep(time.Second)
	cmd := exec.Command("git", "push", "origin", "main")
	cmd.Output()
	fmt.Println("slatt")

}
