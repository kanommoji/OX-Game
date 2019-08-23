package main

import (
	"ox/api"

	"github.com/gin-gonic/gin"
)

func main() {
	var gameApi api.GameAPI
	route := gin.Default()
	route.POST("ox/new-game/player", gameApi.AddPlayer)
	route.POST("ox/new-game", gameApi.SetNewGame)
	route.Run(":8080")
}
