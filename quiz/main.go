package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Problem struct {
	Question string
	Answer   string
}

func createQuizList(data [][]string) []Problem {
	var quizList []Problem
	for _, line := range data {
		var rec Problem
		rec.Question = line[0]
		rec.Answer = line[1]
		quizList = append(quizList, rec)
	}
	return quizList
}

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		fmt.Printf("Failed to open file : %v\n", err)
	}
	defer file.Close()

	r := csv.NewReader(file)

	problems, err := r.ReadAll()
	correctAnswers := 0
	if err != nil {
		fmt.Printf("Failed to reading file : %v\n", err)
	}
	quiz := createQuizList(problems)

	for i, problem := range quiz {
		fmt.Printf("%d. %s?\n", i+1, problem.Question)
		var userAnswer string
		_, err := fmt.Scanln(&userAnswer)
		if err != nil {
			fmt.Printf("Failed to reading user input : %v\n", err)
		}
		if userAnswer == problem.Answer {
			correctAnswers++
		}
	}

	fmt.Printf("Your Answer : %d / %d \n", correctAnswers, len(problems))

}
