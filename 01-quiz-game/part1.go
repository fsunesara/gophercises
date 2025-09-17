package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
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

func Quiz1(fileName string) {
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

	fmt.Printf("You scored %d out of %d.\n", score, numQuestions)
}
