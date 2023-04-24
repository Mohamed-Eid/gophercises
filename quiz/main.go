package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
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
	var (
		filepath string
		duration time.Duration
		shuffle  bool
	)

	flag.StringVar(&filepath, "path", "problems.csv", "file path")
	flag.DurationVar(&duration, "duration", 30*time.Second, "quiz timer")
	flag.BoolVar(&shuffle, "shuffle", false, "shuffle the questions")

	flag.Parse()

	fmt.Printf("exam : %s , duration : %+v , shuffle %v\n", filepath, duration, shuffle)

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Failed to open file : %v\n", err)
	}
	defer file.Close()

	r := csv.NewReader(file)

	problems, err := r.ReadAll()
	if err != nil {
		fmt.Printf("Failed to reading file : %v\n", err)
	}
	quiz := createQuizList(problems)
	correctAnswers := 0
	totalQuestions := len(problems)

	timer := time.NewTimer(duration)

	go func() {
		<-timer.C
		fmt.Println("Time Expired")
		fmt.Printf("Your Answer : %d / %d \n", correctAnswers, totalQuestions)
		os.Exit(0)
	}()

	if shuffle {
		rand.Shuffle(len(quiz), func(i, j int) {
			quiz[i], quiz[j] = quiz[j], quiz[i]
		})
	}

	for i, problem := range quiz {
		fmt.Printf("%d. %s?\n", i+1, problem.Question)
		var userAnswer string
		_, err := fmt.Scanln(&userAnswer)
		if err != nil {
			fmt.Printf("Failed to reading user input : %v\n", err)
		}
		if strings.TrimSpace(userAnswer) == problem.Answer {
			correctAnswers++
		}
	}

	// Stop the timer before it expires
	if !timer.Stop() {
		// Drain the channel to prevent a goroutine leak
		<-timer.C
	}

	fmt.Printf("Your Answer : %d / %d \n", correctAnswers, totalQuestions)

}
