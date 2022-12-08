package csvReader

import (
	"florencesarmah/gophercises/quiz/problem"
    "encoding/csv"
    "log"
    "os"
)

func ReadQuestions(fileName string) []problem.Problem {
	return convertToQuestions(getDataFrom(fileName))
}

func getDataFrom(fileName string) [][]string {
    // open file
    f, err := os.Open("./resources/" + fileName)
    if err != nil {
        log.Fatal(err)
    }

    // remember to close the file at the end of the program
    defer f.Close()

    // read csv values using csv.Reader
    csvReader := csv.NewReader(f)
    data, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal(err)
    }
	return data
}

func convertToQuestions(data [][]string) []problem.Problem {
    var problemList []problem.Problem
    for _, line := range data {
		var problem problem.Problem
		for j, field := range line {
			if j == 0 {
				problem.Question = field
			} else if j == 1 {
				problem.Answer = field
			}
		}
		problemList = append(problemList, problem)
    }
    return problemList
}