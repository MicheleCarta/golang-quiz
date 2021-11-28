package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MicheleCarta/golang-quiz/data"
)

func fetchPlayers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: fetchPlayers")
	var playersResult []Player
	for _, row := range data.DisplayAllPlayers() {
		playersResult = append(playersResult, Player{id: row.id, name: row.username, score, row.score})
		//movies = append(movies, Movie{MovieID: movieID, MovieName: movieName})
	}
	json.NewEncoder(w).Encode(data.DisplayAllPlayers())
}
