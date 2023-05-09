// Part 1
// Create a program that will read in a quiz provided via a CSV file (more details below) and will then give the quiz to a user keeping track of how many questions they get right and how many they get incorrect. Regardless of whether the answer is correct or wrong the next question should be asked immediately afterwards.

// The CSV file should default to problems.csv (example shown below), but the user should be able to customize the filename via a flag.

// The CSV file will be in a format like below, where the first column is a question and the second column in the same row is the answer to that question.

// 5+5,10
// 7+3,10
// 1+1,2
// 8+3,11
// 1+2,3
// 8+6,14
// 3+1,4
// 1+4,5
// 5+1,6
// 2+3,5
// 3+3,6
// 2+4,6
// 5+2,7
// You can assume that quizzes will be relatively short (< 100 questions) and will have single word/number answers.

// At the end of the quiz the program should output the total number of questions correct and how many questions there were in total. Questions given invalid answers are considered incorrect.

// NOTE: CSV files may have questions with commas in them. Eg: "what 2+2, sir?",4 is a valid row in a CSV. I suggest you look into the CSV package in Go and don’t try to write your own CSV parser.

package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Problem struct {
	question string
	answer   string
}

type Result int

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func parseline(lines [][]string) []Problem {
	problems := make([]Problem, len(lines))
	for index, line := range lines {
		problems[index] = Problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return problems
}

func exam(problems []Problem) Result {
	var correctedAnswerCount Result
	for index, problem := range problems {
		var userInputAnswer string

		fmt.Printf("Problem #%d: %v = \n", index+1, problem.question)
		_, err := fmt.Scanf("%s", &userInputAnswer)
		if err != nil {
			exit(fmt.Sprintf("Error converting input: %v \n", err))

		}
		if userInputAnswer == problem.answer {
			correctedAnswerCount++
		}
	}
	return correctedAnswerCount
}

func main() {

	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open file: %s \n", *csvFileName))
	}
	defer file.Close()

	//  Read the CSV data
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Failed to parse file: %s \n", *csvFileName))
	}
	problems := parseline(lines)
	result := exam(problems)
	fmt.Printf("Out of %d you have scored %v \n", len(problems), result)
}
