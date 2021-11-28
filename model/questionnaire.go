package model

import (
	"fmt"
)

// Computer player
type Questionnaire struct {
	Answers []string
	counter int
}

// Print _
func (c *Questionnaire) Print(output string) {
	fmt.Println(output)
}

// Input _
func (q *Questionnaire) Input() (string, error) {
	defer q.incrementCounter()

	return q.Answers[q.counter], nil
}

func (q *Questionnaire) incrementCounter() {
	q.counter++
}
