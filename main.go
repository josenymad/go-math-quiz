package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
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

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for index, question := range questions {
		fmt.Printf("Question %v:\n%v=", index+1, question)
		answerCh := make(chan string)
		go func() {
			var input string
			fmt.Scan(&input)
			answerCh <- input
		}()
		select {
		case <-timer.C:
			for index, userAnswer := range userAnswers {
				if userAnswer == answers[index] {
					correctAnswers++
				} else {
					wrongAnswers = append(wrongAnswers, index+1)
				}
			}
			if len(wrongAnswers) == 1 {
				fmt.Printf("\nTime's up! You got %v out of %v questions right, question %v was incorrect\n", correctAnswers, len(questions), wrongAnswers[0])
			} else if len(wrongAnswers) > 1 {
				fmt.Printf("\nTime's up! You got %v out of %v questions right, these are the questions you got wrong %v\n", correctAnswers, len(questions), wrongAnswers)
			} else {
				fmt.Printf("\nTime's up! All the questions you answered were correct, but you didn't have time to answer the remaining %v questions\n", len(questions)-len(userAnswers))
			}
			return
		case answer := <-answerCh:
			userAnswers = append(userAnswers, answer)
		}
	}

	for index, userAnswer := range userAnswers {
		if userAnswer == answers[index] {
			correctAnswers++
		} else {
			wrongAnswers = append(wrongAnswers, index+1)
		}
	}

	if len(wrongAnswers) == 1 {
		fmt.Printf("\nYou got %v out of %v questions right, question %v was incorrect\n", correctAnswers, len(questions), wrongAnswers[0])
	} else if len(wrongAnswers) > 1 {
		fmt.Printf("\nYou got %v out of %v questions right, these are the questions you got wrong %v\n", correctAnswers, len(questions), wrongAnswers)
	} else {
		fmt.Printf("\nWell done! You got all the questions correct!\n")
	}
}
