package main

import (
	"log"
	"os"
	"os/exec"
	"os/signal"
	"time"
)

func main() {
	cfg := LoadCondig()
	var log log.Logger
	defer func() {
		cmd := exec.Command("git", "push", "origin", "main")
		res, err := cmd.Output()

		if err != nil {
			log.Fatal(err)
		}
		log.Println("exit", res)

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

}
