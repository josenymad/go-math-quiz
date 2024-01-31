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

	fmt.Println(answers, userAnswers)
}
