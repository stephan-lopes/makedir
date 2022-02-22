package main

import (
	"log"

	"github.com/stephan-lopes/makedir/entities"
)

func main() {
	config := entities.NewConfig()

	config, err := config.FromYAML("config.yaml")
	if err != nil {
		log.Fatalf("Error during yaml parse: %v", err)
	}

	work := entities.Worker{
		Path: entities.CreateList(config),
	}

	if err := work.CreateDir(); err != nil {
		log.Fatalf("Error during creation of directories: %v", err)
	}
}
