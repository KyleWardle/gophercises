package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

type QuizQuestion struct {
	Question string
	Answer   string
}

func shuffleQuestions(questions []QuizQuestion) []QuizQuestion {
	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})

	return questions
}

func getQuestionsFromCsv(filename string, shuffle bool) ([]QuizQuestion, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening csv : %w", err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)

	var questions []QuizQuestion

	for {
		row, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				if shuffle {
					questions = shuffleQuestions(questions)
				}

				return questions, nil
			}

			return nil, fmt.Errorf("error reading csv row : %w", err)
		}

		question := QuizQuestion{
			Question: row[0],
			Answer:   row[1],
		}

		questions = append(questions, question)
	}
}

func askQuestion(question QuizQuestion) (bool, error) {
	fmt.Println(question.Question)

	var userAnswer string
	_, err := fmt.Scanln(&userAnswer)
	if err != nil {
		return false, fmt.Errorf("error reading input: %w", err)
	}

	return userAnswer == question.Answer, nil
}

func askQuestions(questions []QuizQuestion, correctQuestions *int) {
	for _, question := range questions {
		correct, err := askQuestion(question)
		if err != nil {
			log.Fatal(err)
		}

		if correct {
			*correctQuestions++
		}
	}
}

func main() {
	filepath := flag.String("filepath", "problems.csv", "filepath of csv for questions")
	allowedTime := flag.Int("time-limit", 30, "Time limit for quiz")
	shuffle := flag.Bool("shuffle", false, "Shuffle quiz questions")
	flag.Parse()

	questions, err := getQuestionsFromCsv(*filepath, *shuffle)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to get questions for quiz : %w", err))
	}

	start, err := askQuestion(QuizQuestion{
		Question: "Type start to begin the quiz!",
		Answer:   "start",
	})
	if err != nil {
		log.Fatal(fmt.Errorf("failed to start quiz : %w", err))
	}

	if start {
		correctQuestions := 0
		go askQuestions(questions, &correctQuestions)

		time.Sleep(time.Duration(*allowedTime*1000) * time.Millisecond)
		fmt.Printf("Your score was %d/%d !\n", correctQuestions, len(questions))
	}

}
