package business

import (
	"flag"
	"log"

	"github.com/MicheleCarta/golang-quiz/game"
	"github.com/MicheleCarta/golang-quiz/game/model"
)

func StartGame() {
	fileName := flag.String("file", "problems.yaml", "The name of the file with the problems")
	limit := flag.Int("limit", 100, "The time limit for the quiz in seconds")
	flag.Parse()

	quiz, err := game.New(*fileName)
	if err != nil {
		log.Fatal(err)
	}

	player := model.Person{}

	gameService := New(*limit, *quiz, &player)
	_, err = gameService.Run()
	if err != nil {
		log.Fatal(err)
	}

}
