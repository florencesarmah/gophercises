//https://github.com/gophercises/quiz

package main

import (
	"florencesarmah/gophercises/quiz/csvReader"
	"florencesarmah/gophercises/quiz/problem"
	"fmt"
	"os"
)

func main() {
	problemList := csvReader.ReadQuestions(getFileName())
	result := problem.StartQuiz(problemList)
	fmt.Println("Got", result.CorrectAnswers, "out of", result.TotalQuestions, "right")
}

func getFileName() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	} else {
		return getDefaultFileName()
	}
}

func getDefaultFileName() string {
	return "problems.csv"
}
