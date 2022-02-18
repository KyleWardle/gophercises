package Adapters

import (
	"choose-your-adventure/Services"
	"fmt"
	template2 "html/template"
	"net/http"
	"strings"
)

type WebAdapter struct {
	service *Services.StoryService
}

func (wa WebAdapter) Initialise(service *Services.StoryService) {
	wa.service = service

	mux := http.NewServeMux()
	mux.HandleFunc("/answer/", wa.setAnswer)
	mux.HandleFunc("/", wa.askQuestion)

	println("Serving on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(fmt.Errorf("error serving handlers : %w", err))
	}

}

func (wa WebAdapter) askQuestion(w http.ResponseWriter, r *http.Request) {
	template, err := template2.ParseFiles("Templates/ask-question.html")
	if err != nil {
		panic(fmt.Errorf("error parsing html file : %w", err))
	}

	err = template.Execute(w, wa.service.GetStoryPart())
	if err != nil {
		panic(fmt.Errorf("error executing html file : %w", err))
	}
}

func (wa WebAdapter) setAnswer(w http.ResponseWriter, r *http.Request) {
	answer := strings.Replace(r.URL.String(), "/answer/", "", -1)
	wa.service.CurrentPart = answer
	http.Redirect(w, r, "/", http.StatusFound)
}
