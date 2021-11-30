package game

import (
	"fmt"
)

type Questionnaire struct {
	Answers []string
	counter int
}

func (c *Questionnaire) Print(output string) {
	fmt.Println(output)
}

func (q *Questionnaire) Input() (string, error) {
	defer q.incrementCounter()

	return q.Answers[q.counter], nil
}

func (q *Questionnaire) incrementCounter() {
	q.counter++
}
