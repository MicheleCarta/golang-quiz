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
	"github.com/MicheleCarta/golang-quiz/cmd"
)

func main() {
	cmd.Execute()
	/**
	data.OpenDatabase()
	data.CreateTablePlayers()
	data.CreateTableQuizScores()
	//data.InsertScore("test", 14.5)
	data.InsertPlayer("user1", 14.5)
	// Init the mux router
	router := mux.NewRouter()
	router.HandleFunc("/players/", getPlayers).Methods("GET")
	router.HandleFunc("/", homePage).Methods("GET")
	fmt.Println("Server at 10000")
	log.Fatal(http.ListenAndServe(":10000", router))

	*/

}
