package main

import (
	"log"

	"github.com/Nepsteren/CLI_TaskTracker.git/cli"
	"github.com/Nepsteren/CLI_TaskTracker.git/service"
	"github.com/Nepsteren/CLI_TaskTracker.git/storage"
)

func main() {
	store := storage.New("tasks.json")

	svc, err := service.New(store)
	if err != nil {
		log.Fatal(err)
	}

	cli.New(svc).Run()
}
