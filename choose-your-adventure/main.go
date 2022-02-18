package main

import (
	"choose-your-adventure/Adapters"
	"choose-your-adventure/Services"
	"flag"
	"fmt"
)

type DisplayAdapter interface {
	Initialise(service *Services.StoryService)
}

func main() {
	adapter := flag.String("adapter", "console", "Adapter to run the application in")
	flag.Parse()

	service := Services.StoryService{}
	err := service.SetupStoryFromJson("story.json")
	if err != nil {
		panic(fmt.Errorf("error getting story : %w", err))
	}

	displayAdapter := getDisplayAdapter(*adapter)
	displayAdapter.Initialise(&service)
}

func getDisplayAdapter(adapter string) DisplayAdapter {
	if adapter == "web" {
		return Adapters.WebAdapter{}
	} else {
		return Adapters.ConsoleAdapter{}
	}
}
