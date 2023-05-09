// Adapt your program from part 1 to add a timer. The default time limit should be 30 seconds, but should also be customizable via a flag.

// Your quiz should stop as soon as the time limit has exceeded. That is, you shouldn’t wait for the user to answer one final questions but should ideally stop the quiz entirely even if you are currently waiting on an answer from the end user.

// Users should be asked to press enter (or some other key) before the timer starts, and then the questions should be printed out to the screen one at a time until the user provides an answer. Regardless of whether the answer is correct or wrong the next question should be asked.

// At the end of the quiz the program should still output the total number of questions correct and how many questions there were in total. Questions given invalid answers or unanswered are considered incorrect.

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

func exam(problems []Problem, timeLimit *int) Result {

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	var correctedAnswerCount Result
	for index, problem := range problems {
		fmt.Printf("Problem #%d: %v = \n", index+1, problem.question)
		answerChannel := make(chan string)

		// Go routin for non blocking call for timer
		go func() {
			var userInputAnswer string
			fmt.Scanf("%s", &userInputAnswer)
			answerChannel <- userInputAnswer
		}()

		select {
		case <-timer.C:
			return correctedAnswerCount
		case answer := <-answerChannel:
			if answer == problem.answer {
				correctedAnswerCount++
			}
		}
	}
	return correctedAnswerCount
}

func shuffle(lines [][]string) [][]string {
	rand.Shuffle(len(lines), func(i, j int) {
		lines[i], lines[j] = lines[j], lines[i]
	})
	return lines
}

func main() {

	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("timeLimit", 30, "Enter the duration for quiz")

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
	shuffleLines := shuffle(lines)

	problems := parseline(shuffleLines)

	result := exam(problems, timeLimit)
	fmt.Printf("Out of %d you have scored %v \n", len(problems), result)
}
