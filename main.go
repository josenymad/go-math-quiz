package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	questions := []string{}
	answers := []string{}
	userAnswers := []string{}
	correctAnswers := 0
	wrongAnswers := []int{}

	for _, questionAnswer := range records {
		questions = append(questions, questionAnswer[0])
		answers = append(answers, questionAnswer[1])
	}

	for _, question := range questions {
		var input string
		fmt.Println(question)
		fmt.Scan(&input)
		userAnswers = append(userAnswers, input)
	}

	for index, userAnswer := range userAnswers {
		if userAnswer == answers[index] {
			correctAnswers++
		} else {
			wrongAnswers = append(wrongAnswers, index+1)
		}
	}

	fmt.Printf("You got %v out of %v questions right, these are the questions you got wrong %v", correctAnswers, len(questions), wrongAnswers)
}
