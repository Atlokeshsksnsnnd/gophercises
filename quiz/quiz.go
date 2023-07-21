package main

import (
	"fmt"
	"time"
)

type problem struct {
	question string
	answer   string
}

type quiz struct {
	problems []problem
	mark     int
}

func (q *quiz) parseProblems(problems [][]string) {
	q.problems = make([]problem, len(problems))
	for i, line := range problems {
		p := problem{question: line[0], answer: line[1]}
		q.problems[i] = p
	}
}

func (q *quiz) startExam(timeLimit int) {
	status := make(chan bool)
	go q.askQuestion(status)
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	for {
		select {
		case <-status:
			return
		case <-timer.C:
			return
		}
	}
}

func (q *quiz) askQuestion(status chan bool) {
	var answer string
	for _, p := range q.problems {
		fmt.Println(p.question)
		fmt.Scan(&answer)
		if p.answer == answer {
			q.mark++
		}
	}
	status <- true
}

func (q *quiz) printResult() {
	fmt.Println("Total Marks ", len(q.problems), " Marks Obtained ", q.mark)
}
