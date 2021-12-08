package business

import (
	"fmt"
	"log"
	"math"
	"time"

	"encoding/json"

	"github.com/MicheleCarta/golang-quiz/data"
	"github.com/MicheleCarta/golang-quiz/game"
	"github.com/MicheleCarta/golang-quiz/game/model"
	"github.com/MicheleCarta/golang-quiz/pkg/websocket"
	"github.com/MicheleCarta/golang-quiz/service"
)

var (
	games    = make(map[float64]*ServiceGame)
	timer    = time.NewTimer(time.Duration(100) * time.Second)
	clientWS = make(map[float64]*websocket.Client)
)

func AddWsClientConnection(client *websocket.Client) {
	clientWS[client.ID] = client
}

func NewWeb(limit int, quiz game.Quiz, player Player, idPlayer float64, currentScore int, gameMatch int) *ServiceGame {
	return &ServiceGame{limit, quiz, player, idPlayer, currentScore, 0, gameMatch}
}

func SubscribeGame(file string, playerId float64) *ServiceGame {
	quiz, _ := game.New(file)
	var playerEntity = data.GetPlayer(playerId)
	player := model.Person{}
	var gameService *ServiceGame = NewWeb(100, *quiz, &player, playerEntity.Id, int(playerEntity.Score), service.GetPlayer(playerId).GameMatch)
	games[playerEntity.Id] = gameService
	log.Println("added game in memory : ", games[playerEntity.Id].quiz)
	return games[playerId]
}

func Round(playerId float64, answer string) bool {
	var res bool = true
	var s *ServiceGame = games[playerId]
	var pos = 0
	if s.quiz.Problems[pos].Correct == answer {
		service.InsertScore(s.idPlayer, s.quiz.Problems[pos].Question, true)
		games[playerId].score++
	} else {
		service.InsertScore(s.idPlayer, s.quiz.Problems[pos].Question, false)
		res = false
	}
	games[playerId].quiz.Problems = append(s.quiz.Problems[:pos], s.quiz.Problems[pos+1:]...)
	fmt.Println("current len player :", playerId, " Problems ", len(games[playerId].quiz.Problems))
	//body := json.Marshal(games[playerId].quiz.Problems[0])

	if len(games[playerId].quiz.Problems) <= 0 {
		var players []data.Player = service.FetchPlayers()
		var totalScores = make([]int, 0, len(players))
		for _, ps := range players {
			totalScores = append(totalScores, int(ps.Score))

		}
		var percentage = math.Round(((float64(games[playerId].currentScore+games[playerId].score) / game.FindAverage(totalScores) / 100) * 1000))
		fmt.Println("You were better than ", percentage, "% of all quizzers ")
		service.UpdatePlayer(playerId, games[playerId].score, games[playerId].currentScore, percentage, (s.gameMatch + 1))
		body := fmt.Sprintf("%s%f%s", "Game finished You were better than ", percentage, "% of all quizzers ")
		if err := clientWS[playerId].Conn.WriteJSON(websocket.Message{Type: 2, Body: body}); err != nil {
			fmt.Println(err)
		}
	} else {
		body, err := json.Marshal(games[playerId].quiz.Problems[0])
		if err != nil {
			fmt.Println(err)
		}
		if err := clientWS[playerId].Conn.WriteJSON(websocket.Message{Type: 1, Body: string(body)}); err != nil {
			fmt.Println(err)
		}
	}
	return res
}

func CurrentGame(playerId float64) game.Quiz {
	var quiz = games[playerId].quiz
	var problems []game.Problem
	for _, pr := range quiz.Problems {
		pr.Correct = "***"
		problems = append(problems, pr)
	}
	quiz.Problems = problems
	return quiz
}
