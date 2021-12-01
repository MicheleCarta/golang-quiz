package service

import (
	"github.com/MicheleCarta/golang-quiz/data"
)

func FetchPlayers() []data.Player {
	return data.DisplayAllPlayers()
}
func AddPlayer(name string, score float64) {
	data.InsertPlayer(name, score)
}
func UpdatePlayer(score int, idPlayer float64) {
	data.UpdatePlayer(score, idPlayer)
}

func GetPlayer(playerId float64) data.Player {
	return data.GetPlayer(playerId)
}

func InsertScore(idPlayer float64, question string, outcome bool) {
	data.InsertScore(idPlayer, question, outcome)
}

func GetScoresPlayer(idPlayer float64) []data.Scores {
	return data.GetScoresPlayer(idPlayer)
}
