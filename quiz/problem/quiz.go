package problem

import (
    "fmt"
)

type Problem struct {
    Question	string
    Answer		string
}

type QuizResult struct {
    CorrectAnswers int
    TotalQuestions int
}

func StartQuiz(problemList []Problem) QuizResult {
    quizResult := quizUser(problemList)
    quizResult.TotalQuestions = len(problemList)
    return quizResult
}

func quizUser(problemList []Problem) QuizResult {
	var quizResult QuizResult
	for _, problem := range problemList {
		fmt.Print(problem.Question, ": ")
		var userAnswer string
		fmt.Scanln(&userAnswer)
		if problem.Answer == userAnswer {
			quizResult.CorrectAnswers += 1
		}
	}
	return quizResult
}