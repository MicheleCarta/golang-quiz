package business

import (
	"fmt"
	"strings"
	"time"

	"github.com/MicheleCarta/golang-quiz/game"
	"github.com/MicheleCarta/golang-quiz/service"
)

// Player interface
type Player interface {
	Print(string)
	Input() (string, error)
}

type serviceGame struct {
	limit  int
	quiz   game.Quiz
	player Player
}

// New _
func New(limit int, quiz game.Quiz, player Player) service.Service {
	return &serviceGame{limit, quiz, player}
}
func (s *serviceGame) Run() (int, error) {
	t := time.NewTimer(time.Duration(s.limit) * time.Second)

	var score int
	for i, prob := range s.quiz.Problems {
		s.player.Print(fmt.Sprintf("Problem %d: %s ?", i+1, prob.Question))
		for j, ans := range prob.Answer {
			s.player.Print(fmt.Sprintf("Answer %d: %s ?", j+1, ans))
		}
		type Answer struct {
			input string
			err   error
		}

		answerChan := make(chan Answer)

		go func() {
			input, err := s.player.Input()
			if err != nil {
				answerChan <- Answer{
					err: err,
				}
			}

			input = strings.TrimSpace(input)

			answerChan <- Answer{
				input: input,
			}
		}()

		select {
		case <-t.C:
			goto GameOver
		case answer := <-answerChan:
			if answer.err != nil {
				return score, answer.err
			}

			if answer.input == prob.Correct {
				s.player.Print("Correct\n")
				service.InsertScore(1, prob.Question, true)
				score++
			} else {
				s.player.Print("Wrong\n")
				service.InsertScore(1, prob.Question, false)
			}
		}
	}

GameOver:
	s.player.Print(fmt.Sprintf("Game over! Your score is %d from %d", score, len(s.quiz.Problems)))

	return score, nil
}
