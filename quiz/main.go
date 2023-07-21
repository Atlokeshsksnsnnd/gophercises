package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {

	// Parses the commmand line arguments
	fileName := flag.String("filename", "problem.csv", "CSV file containg quiz")
	timeLimit := flag.Int("timelimit", 30, "TimeLimit for the Quiz")
	flag.Parse()

	// Opens a file
	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Reads the csv file
	csvReader := csv.NewReader(file)
	problems, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Struct containg the problems
	quiz := quiz{}
	quiz.parseProblems(problems)
	quiz.startExam(*timeLimit)

	quiz.printResult()
}
