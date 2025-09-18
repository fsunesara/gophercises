package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func readCsv(fileName string) [][]string {
	f, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(strings.NewReader(string(f)))

	questions, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	return questions
}

func quiz1(fileName string) {
	questions := readCsv(fileName)
	score := 0
	numQuestions := len(questions)
	for _, question := range questions {
		var answer string
		fmt.Printf("%s = ", question[0])
		fmt.Scanln(&answer)
		if answer == question[1] {
			score++
		}
	}

	endQuiz(score, numQuestions)
}

func quiz2(fileName string, timeLimit int) {
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	questions := readCsv(fileName)
	score := 0
	numQuestions := len(questions)

	go func() {
		<-timer.C
		fmt.Println("\ntimeout")
		endQuiz(score, numQuestions)
	}()

	for _, question := range questions {
		var answer string
		fmt.Printf("%s = ", question[0])
		fmt.Scanln(&answer)
		if answer == question[1] {
			score++
		}
	}

	endQuiz(score, numQuestions)
}

func endQuiz(score int, numQuestions int) {
	fmt.Printf("You scored %d out of %d.\n", score, numQuestions)
	os.Exit(0)
}

func main() {
	fileName := flag.String("file", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("timeLimit", 30, "the time limit for the quiz in seconds")
	flag.Parse()
	// quiz1(*fileName)
	quiz2(*fileName, *timeLimit)
}
