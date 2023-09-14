package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"time"
)

func main() {
	cfg := LoadCondig()
	var log log.Logger
	file, _ := os.OpenFile(cfg.FileToReadName, os.O_RDWR|os.O_CREATE, 0666)
	defer file.Close()
	defer func() {
		cmd := exec.Command("git", "push", "origin", "main")
		_, _ = cmd.Output()

	}()
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)

	gitinit(&cfg)
	go func() {
		for {
			time.Sleep(time.Second)

			file, err := createFile()
			if err != nil {
				log.Print(err)

				os.Exit(1)
			}

			if err := add(file); err != nil {
				log.Print(err)

				os.Exit(1)
			}
			if err := commit(file); err != nil {
				log.Print(err)

				os.Exit(1)
			}
			if err := deleteFile(file); err != nil {
				log.Print(err)
				os.Exit(1)

			}

		}
	}()
	<-exit
	fmt.Println("exit")

}
