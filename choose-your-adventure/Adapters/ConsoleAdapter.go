package Adapters

import (
	"choose-your-adventure/Services"
	"fmt"
	"time"
)

type ConsoleAdapter struct {
	service *Services.StoryService
}

func (ca ConsoleAdapter) Initialise(service *Services.StoryService) {
	ca.service = service
	ca.askQuestion()
}

func (ca ConsoleAdapter) askQuestion() {
	storyPart := ca.service.GetStoryPart()

	storyLines := storyPart.Story
	for _, line := range storyLines {
		fmt.Println(line)
		fmt.Println("")
		time.Sleep(ca.getConsoleWaitTime(line))
	}

	if ca.service.StoryHasFinished() {
		fmt.Println("End of Story!")
		return
	}

	for key, option := range storyPart.Options {
		fmt.Println(fmt.Sprintf("Enter %d to: %s", key, option.Text))
	}

	var text string
	fmt.Scanln(&text)
	ca.service.SelectStoryOption(text)

	fmt.Println("\n\n\n")
	ca.askQuestion()
}

func (ca ConsoleAdapter) getConsoleWaitTime(text string) time.Duration {
	//wpm := 300      // Readable words per minute
	wpm := 3000     // Readable words per minute
	wordLength := 6 // average word length
	wordCount := len(text) / wordLength
	timeToRead := ((float32(wordCount) / float32(wpm)) * float32(60)) * float32(1000)

	return time.Duration(timeToRead) * time.Millisecond
}
