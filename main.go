package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFileName))
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		exit("Could not parse the provided CSV file")
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

	for index, question := range questions {
		var input string
		fmt.Printf("Question %v:\n%v=", index+1, question)
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

	fmt.Printf("You got %v out of %v questions right, these are the questions you got wrong %v\n", correctAnswers, len(questions), wrongAnswers)
}
