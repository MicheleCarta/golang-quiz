package business

import (
	"flag"
	"fmt"
	"log"

	"github.com/MicheleCarta/golang-quiz/game"
	"github.com/MicheleCarta/golang-quiz/game/model"
	"github.com/MicheleCarta/golang-quiz/service"
)

func StartGame() {
	fileName := flag.String("file", "problems.yaml", "The name of the file with the problems")
	limit := flag.Int("limit", 100, "The time limit for the quiz in seconds")
	flag.Parse()

	quiz, err := game.New(*fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Choice the Player \n")
	for i, pl := range service.FetchPlayers() {
		fmt.Printf("Select Player - just choice the number %d: %s ?", i, " ", pl.Username, pl.Id)
	}
	var id float64
	fmt.Scanln(&id)
	player := model.Person{}

	gameService := New(*limit, *quiz, &player, id)
	_, err = gameService.Run()
	if err != nil {
		log.Fatal(err)
	}

}
