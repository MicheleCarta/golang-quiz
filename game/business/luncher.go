package business

import (
	"flag"
	"fmt"
	"log"

	"github.com/MicheleCarta/golang-quiz/data"
	"github.com/MicheleCarta/golang-quiz/game"
	"github.com/MicheleCarta/golang-quiz/game/model"
	"github.com/MicheleCarta/golang-quiz/service"
)

/**Select a Player and Play*/
func StartGame() {
	var players []data.Player = service.FetchPlayers()
	id, _ := choicePlayer(players)
	player := model.Person{}

	quiz, limit, err := initGame()
	gameService := New(limit, *quiz, &player, id)
	_, err = gameService.Run()
	if err != nil {
		log.Fatal(err)
	}
	var playerscores []data.Player = service.FetchPlayers()
	var totalScores = make([]int, 0, len(playerscores))
	for _, ps := range playerscores {
		totalScores = append(totalScores, int(ps.Score))

	}
	_, min, max := findMinAndMax(totalScores)
	fmt.Println("You were better than ", PercentageChange(min, max))
}

/** it will consume all the Player*/
func StartAutoGame() {
	scores := make(map[string]int)
	totalScores := make([]int, 0, len(scores))
	playersScore := make([]string, 0, len(scores))
	quiz, limit, err := initGame()
	var players []data.Player = service.FetchPlayers()
	player := model.Person{}
	for _, pl := range players {
		fmt.Println(pl.Username, " is the current Player")
		gameService := New(limit, *quiz, &player, pl.Id)
		var score = 0
		score, err = gameService.Run()
		scores[pl.Username] = score
		if err != nil {
			log.Fatal(err)
		}
	}
	for username, score := range scores {
		fmt.Println("scores ", username, score)
		totalScores = append(totalScores, score)
		playersScore = append(playersScore, username)

	}
	_, min, max := findMinAndMax(totalScores)
	fmt.Println("the max and min scores are ", min, max)
	fmt.Println("You were better than ", PercentageChange(min, max))
}

func initGame() (*game.Quiz, int, error) {
	fileName := flag.String("file", "problems.yaml", "The name of the file with the problems")
	limit := flag.Int("limit", 100, "The time limit for the quiz in seconds")
	flag.Parse()
	quiz, err := game.New(*fileName)
	if err != nil {
		log.Fatal(err)
	}
	return quiz, *limit, err
}

func choicePlayer(players []data.Player) (float64, int) {
	fmt.Printf("Choice the Player \n")
	var i = 0
	for i, pl := range players {
		fmt.Printf("[%v]: %s  %d \n", pl.Id, pl.Username, i)
	}
	var id float64
	fmt.Scanln(&id)
	return id, i
}

func PercentageChange(old, new int) (delta float64) {
	diff := float64(new - old)
	delta = (diff / float64(old)) * 100
	return
}

func findMinAndMax(scores []int) (min int, max int, index int) {
	min = scores[0]
	max = scores[0]
	var j = 0
	for i, value := range scores {
		if value < min {
			min = value
		}
		if value > max {
			max = value
			j = i
		}
	}
	return min, max, j
}
