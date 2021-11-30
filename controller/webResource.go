package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MicheleCarta/golang-quiz/game/business"
	"github.com/MicheleCarta/golang-quiz/service"
	"github.com/gorilla/mux"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: homePage")
}

func AddPlayer(w http.ResponseWriter, r *http.Request) {
	service.AddPlayer(r.FormValue("username"), 0)
}

func StartGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the best quiz!")
	business.StartGame()
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
