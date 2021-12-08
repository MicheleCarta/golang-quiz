package controller

import (
	"strconv"

	"github.com/MicheleCarta/golang-quiz/game"
	"github.com/MicheleCarta/golang-quiz/game/business"
	"github.com/MicheleCarta/golang-quiz/service"
	"github.com/gin-gonic/gin"
)

func Subscribe(c *gin.Context) {
	if playerId, err := strconv.ParseFloat(c.Param("playerId"), 64); err == nil {
		business.SubscribeGame("problems.yaml", playerId)
		c.JSON(200, gin.H{
			"Subscribed ": playerId,
		})
	}
}

func AddPlayer(c *gin.Context) {
	service.AddPlayer(c.Query("username"), 0, 0.0)
	c.JSON(200, gin.H{
		"Player added": c.Query("username"),
	})
}

func GetPlayers(c *gin.Context) {
	c.JSON(200, gin.H{
		"Players": service.FetchPlayers(),
	})
}

func StartGame(c *gin.Context) {
	res := business.StartGame(false)
	c.JSON(200, gin.H{
		"Quist Started": res,
	})
}

func GetPlayer(c *gin.Context) {
	if playerId, err := strconv.ParseFloat(c.Param("playerId"), 64); err == nil {
		c.JSON(200, gin.H{
			"player": service.GetPlayer(playerId),
		})
	}
}

func GetScoresPlayer(c *gin.Context) {
	if playerId, err := strconv.ParseFloat(c.Param("playerId"), 64); err == nil {
		c.JSON(200, gin.H{
			"score": service.GetScoresPlayer(playerId),
		})
	}
}

func ShowProblems(c *gin.Context) {
	c.JSON(200, gin.H{
		"Problems": service.GetQuizProblems("problems.yaml").Problems,
	})
}
func GetProblems(c *gin.Context) {
	if playerId, err := strconv.ParseFloat(c.Param("playerId"), 64); err == nil {
		var s game.Quiz = business.CurrentGame(playerId)
		c.JSON(200, gin.H{
			"Problems": s.Problems[0],
		})
	}
}

func SendAnswer(c *gin.Context) {
	if playerId, err := strconv.ParseFloat(c.Param("playerId"), 64); err == nil {
		c.JSON(200, gin.H{
			"answer": business.Round(playerId, c.Query("answer")),
		})

	}
}
