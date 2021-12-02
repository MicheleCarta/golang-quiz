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
package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MicheleCarta/golang-quiz/controller"
	"github.com/MicheleCarta/golang-quiz/data"
	"github.com/MicheleCarta/golang-quiz/game/business"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ChoiceAction()
	},
}

var (
	router       = mux.NewRouter()
	isAgain bool = false
)

func init() {

	rootCmd.AddCommand(initCmd)
	data.OpenDatabase()
	ChoiceAction()
}
func ChoiceAction() {
	var play = "p"
	var run = "r"
	var initdb = "i"
	var exit = "e"

repeatAction:
	fmt.Printf("  Choice your next step: \n  [%s] [%s] [%s] [%s] \n  play \n  run server  \n  init DB \n  exit  ", play, run, initdb, exit)
	var ans string
	fmt.Scanln(&ans)
	if ans == play {
		isAgain = business.StartGame(false)
		if isAgain {
			goto repeatAction
		}
	} else if ans == run {
		initApi()
	} else if ans == initdb {
		initDB()
	} else {
		os.Exit(3)
	}
}

func initApi() {
	router.HandleFunc("/", controller.HomePage).Methods("GET")
	router.HandleFunc("/addPlayer/", controller.AddPlayer).Methods("POST")
	router.HandleFunc("/play/", controller.StartGame).Methods("GET")
	router.HandleFunc("/score/{playerId}", controller.GetScoresPlayer).Methods("GET")
	router.HandleFunc("/player/{playerId}", controller.GetPlayer).Methods("GET")
	router.HandleFunc("/players/", controller.GetPlayers).Methods("GET")
	log.Fatal(http.ListenAndServe(":10000", router))
	fmt.Println("Server at 10000")
}
func initDB() {
	data.DropTablePlayer()
	data.DropTableQuizScore()
	data.CreateTablePlayers()
	data.CreateTableQuizScores()
	data.InsertPlayer("Guest", 0, 0.0)
	data.InsertPlayer("Mix", 0, 0.0)
	data.InsertPlayer("Zena", 0, 0.0)
}
