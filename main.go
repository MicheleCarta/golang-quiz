/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MicheleCarta/golang-quiz/cmd"
	"github.com/MicheleCarta/golang-quiz/data"
)

func main() {
	cmd.Execute()
	data.OpenDatabase()
	data.CreateTablePlayers()
	data.CreateTableQuizScores()
	data.InsertScore("test", 14.5)
	data.DisplayAllPlayers()
	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/players", service.fetchPlayers())
	log.Fatal(http.ListenAndServe(":10000", nil))
}
