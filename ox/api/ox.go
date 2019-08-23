package api

import (
	"net/http"
	service "ox/service"

	"github.com/gin-gonic/gin"
)

type GameAPI struct {
	GameAPI service.Game
}

func (gameAPI *GameAPI) AddPlayer(c *gin.Context) {
	var player service.Player
	err := c.BindJSON(&player)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	gameAPI.GameAPI.Player = append(gameAPI.GameAPI.Player, player)
}

func (gameAPI *GameAPI) SetNewGame(c *gin.Context) {
	gameAPI.GameAPI = service.NewGame(gameAPI.GameAPI.Player, "x")
}
