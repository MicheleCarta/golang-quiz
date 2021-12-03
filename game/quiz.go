package game

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Quiz struct {
	Problems []Problem `yaml:"problems"`
}

type Problem struct {
	Question string   `yaml:"question"`
	Answer   []string `yaml:"answer"`
	Correct  string   `yaml:"correct"`
}

func (q Quiz) GetProblems() []Problem {
	return q.Problems
}

type QuizGame struct {
	Quiz
}

func New(fileName string) (*Quiz, error) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return &Quiz{}, err
	}
	var quiz Quiz
	err = yaml.Unmarshal(b, &quiz)
	if err != nil {
		return &Quiz{}, err
	}

	return &quiz, nil
}
