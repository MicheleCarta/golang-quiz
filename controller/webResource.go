package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MicheleCarta/golang-quiz/game"
	"github.com/MicheleCarta/golang-quiz/game/business"
	"github.com/MicheleCarta/golang-quiz/service"
	"github.com/gorilla/mux"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	if playerId, err := strconv.ParseFloat(r.FormValue("playerId"), 64); err == nil {
		fmt.Fprintf(w, "Welcome to the best quiz!")
		business.SubscribeGame("problems.yaml", playerId)
	}

}

func AddPlayer(w http.ResponseWriter, r *http.Request) {
	service.AddPlayer(r.FormValue("username"), 0, 0.0)
}

func GetPlayers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(service.FetchPlayers())

}

func StartGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the best quiz!")
	business.StartGame(false)
	//business.PlayAgain()
	//business.StartAutoGame()
}

func GetPlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	playerID := params["playerId"]
	if playerId, err := strconv.ParseFloat(playerID, 64); err == nil {
		json.NewEncoder(w).Encode(service.GetPlayer(playerId))
	}

}

func GetScoresPlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	playerID := params["playerId"]
	if playerId, err := strconv.ParseFloat(playerID, 64); err == nil {
		json.NewEncoder(w).Encode(service.GetScoresPlayer(playerId))

	}
}

func ShowProblems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(service.GetQuizProblems("problems.yaml").Problems)
}
func GetProblems(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	playerID := params["playerId"]
	if playerId, err := strconv.ParseFloat(playerID, 64); err == nil {
		var s game.Quiz = business.CurrentGame(playerId)
		json.NewEncoder(w).Encode(s.Problems)
	}
}

func SendAnswer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	playerID := params["playerId"]
	var answer = r.FormValue("answer")
	if playerId, err := strconv.ParseFloat(playerID, 64); err == nil {
		json.NewEncoder(w).Encode(business.Round(playerId, answer))

	}
}
