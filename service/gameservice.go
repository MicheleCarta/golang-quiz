package service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MicheleCarta/golang-quiz/data"
)

func FetchPlayers(w http.ResponseWriter, r *http.Request) {
	var response = data.JsonResponse{Type: "success", Data: data.DisplayAllPlayers()}
	json.NewEncoder(w).Encode(response)
}
func AddPlayer(name string, score float64) {
	data.InsertPlayer(name, score)
}

func GetPlayer(playerId float64) data.Player {
	return data.GetPlayer(playerId)
}

func InsertScore(idPlayer float64, question string, outcome bool) {
	log.Println("InsertScore --> ", idPlayer, " ", question, " ", outcome)
	data.InsertScore(1, "question", false)
}

func GetScoresPlayer(idPlayer float64) []data.Scores {
	return data.GetScoresPlayer(idPlayer)
}
