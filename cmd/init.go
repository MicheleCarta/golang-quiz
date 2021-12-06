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

	"github.com/gin-gonic/gin"

	"github.com/MicheleCarta/golang-quiz/controller"
	"github.com/MicheleCarta/golang-quiz/data"
	"github.com/MicheleCarta/golang-quiz/game/business"
	"github.com/MicheleCarta/golang-quiz/pkg/websocket"
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
	fmt.Printf("  Choice your next step:  \n  play \n  run server  \n  init DB \n  exit \n  [%s] [%s] [%s] [%s] ", play, run, initdb, exit)
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
	var r = gin.New()
	r.GET("/home/", func(c *gin.Context) {
		controller.GetProblems(c)
	})

	r.GET("/subscribe/:playerId", func(c *gin.Context) {
		controller.Subscribe(c)
	})
	r.POST("/addPlayer/", func(c *gin.Context) {
		controller.AddPlayer(c)
	})
	r.GET("/play/", func(c *gin.Context) {
		controller.StartGame(c)
	})
	r.GET("/score/:playerId", func(c *gin.Context) {
		controller.GetScoresPlayer(c)
	})
	r.GET("/player/:playerId", func(c *gin.Context) {
		controller.GetPlayer(c)
	})
	r.GET("/problems/:playerId", func(c *gin.Context) {
		controller.GetProblems(c)
	})
	r.POST("/sendAnswer/:playerId", func(c *gin.Context) {
		controller.SendAnswer(c)
	})
	r.GET("/problems/", func(c *gin.Context) {
		controller.ShowProblems(c)
	})
	r.GET("/players/", func(c *gin.Context) {
		controller.GetPlayers(c)
	})
	/**Websocket init */
	pool := websocket.NewPool()
	go pool.Start()

	r.GET("/ws", func(c *gin.Context) {
		fmt.Println("connected! ")
		serveWs(pool, c)
	})
	fmt.Println("Server at 3000")

	if err := r.Run(":3000"); err != nil {
		log.Fatal(err)
	}

}

func HealthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}
	c.JSON(http.StatusOK, res)
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

func serveWs(pool *websocket.Pool, c *gin.Context) {
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := websocket.Upgrade(c)
	if err != nil {
		fmt.Fprintf(c.Writer, "%+v\n", err)
	}

	client := &websocket.Client{
		ID:   0,
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}
