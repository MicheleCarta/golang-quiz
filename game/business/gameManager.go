package business

import (
	"flag"
	"fmt"
	"log"
	"math"

	"github.com/MicheleCarta/golang-quiz/data"
	"github.com/MicheleCarta/golang-quiz/game"
	"github.com/MicheleCarta/golang-quiz/game/model"
	"github.com/MicheleCarta/golang-quiz/service"
)

var (
	fileName *string
	limit    *int
	quiz     *game.Quiz
)

func PlayAgain() bool {
	fmt.Printf("Play Again , y,n ? \n")
	var ans string
	var res = false
	fmt.Scanln(&ans)
	if ans == "y" {
		StartGame(true)
		res = true
	} else {
		fmt.Printf("Thanks to Play")
		//os.Exit(3)
	}
	return res
}

/**Select a Player and Play*/
func StartGame(playAgain bool) bool {
	var players []data.Player = service.FetchPlayers()
	id, _, currentScore, match := choicePlayer(players)
	player := model.Person{}
	var gameService service.Service
	if !playAgain {
		quiz, limit, _ := initGame()
		gameService = New(*&limit, *quiz, &player, id, currentScore, match)
	} else {
		fileName := "problems.yaml"
		limit := 100
		quiz, err := game.New(fileName)
		if err != nil {
			log.Fatal(err)
		}
		gameService = New(limit, *quiz, &player, id, currentScore, match)
	}
	score, err := gameService.Run()
	if err != nil {
		log.Fatal(err)
	}
	var totalScores = make([]int, 0, len(players))
	for _, ps := range players {
		totalScores = append(totalScores, int(ps.Score))

	}
	_, min, max := game.FindMinAndMax(totalScores)
	var percentage = math.Round(((float64(score+currentScore) / game.FindAverage(totalScores) / 100) * 1000))
	fmt.Println("the max and min scores are ", min, max)
	fmt.Println("You were better than ", percentage, "% of all quizzers ")
	service.UpdatePlayer(id, score, currentScore, percentage, (match + 1))
	return PlayAgain()
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
		gameService := New(limit, *quiz, &player, pl.Id, int(pl.Score), pl.GameMatch)
		var score = 0
		score, err = gameService.Run()
		scores[pl.Username] = score
		if err != nil {
			log.Fatal(err)
		}
		service.UpdatePlayer(pl.Id, score, int(pl.Score), pl.Percentage, (pl.GameMatch + 1))
	}
	for username, score := range scores {
		fmt.Println("scores ", username, score)
		totalScores = append(totalScores, score)
		playersScore = append(playersScore, username)

	}
	_, min, max := game.FindMinAndMax(totalScores)
	fmt.Println("the max and min scores are ", min, max)
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

func choicePlayer(players []data.Player) (float64, int, int, int) {
	fmt.Printf("Choice the Player \n")
	var i = 0
	for i, pl := range players {
		fmt.Printf("[%v]: %s %v %d  \n", pl.Id, pl.Username, pl.Percentage, i)
	}
	var id float64
	fmt.Scanln(&id)
	currentPlayer := service.GetPlayer(id)
	score := int(currentPlayer.Score)
	match := int(currentPlayer.GameMatch)
	return id, i, score, match
}
