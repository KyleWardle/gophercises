package Services

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"os"
	"strconv"
)

type StoryService struct {
	Story       map[string]StoryPart
	CurrentPart string
}

type StoryPart struct {
	Title   string
	Story   []string
	Options []StoryOption
}

type StoryOption struct {
	Text string
	Arc  string
}

func (service *StoryService) SetupStoryFromJson(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("error opening json file : %w", err)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading json file : %w", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return fmt.Errorf("error Unmarshalling json file to map : %w", err)
	}

	storyParts := map[string]StoryPart{}
	for key, story := range result {
		var part StoryPart
		err = mapstructure.Decode(story, &part)
		if err != nil {
			return fmt.Errorf("error decoding map file to struct : %w", err)
		}

		storyParts[key] = part
	}

	service.Story = storyParts

	return nil
}

func (service *StoryService) GetStoryPart() StoryPart {
	if service.CurrentPart == "" {
		service.CurrentPart = "intro"
	}

	return service.Story[service.CurrentPart]
}

func (service *StoryService) SelectStoryOption(index string) {
	optionIndex, _ := strconv.Atoi(index)

	part := service.GetStoryPart()
	if len(part.Options)-1 >= optionIndex {
		// if has index
		selectedAnswer := part.Options[optionIndex]
		service.CurrentPart = selectedAnswer.Arc
	}
	// If they give an invalid response, we just throw them around the same way again
}

func (service *StoryService) StoryHasFinished() bool {
	part := service.GetStoryPart()
	return !(len(part.Options) > 0)
}
