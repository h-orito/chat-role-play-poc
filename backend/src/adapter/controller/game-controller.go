package controller

import (
	application "chat-role-play-poc/src/application/service"
	model "chat-role-play-poc/src/domain/model/game"

	"github.com/gin-gonic/gin"
)

type GameController struct {
	gameService application.GameService
}

func NewGameController(gameService application.GameService) GameController {
	return GameController{
		gameService: gameService,
	}
}

type getRequest struct {
	ID uint32 `uri:"id" binding:"required"`
}

func (controller *GameController) Get(c *gin.Context) {
	var getRequest getRequest
	if err := c.ShouldBindUri(&getRequest); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	game, err := controller.gameService.Find(getRequest.ID)
	if err != nil {
		c.JSON(500, gin.H{"msg": err})
		return
	}

	c.JSON(200, gin.H{
		"game": game,
	})
}

type postRequest struct {
	Name string `json:"name" binding:"required"`
}

func (controller *GameController) Post(c *gin.Context) {
	var postRequest postRequest
	if err := c.ShouldBindJSON(&postRequest); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	game := model.Game{
		Name: postRequest.Name,
	}
	saved, err := controller.gameService.Save(game)
	if err != nil {
		c.JSON(500, gin.H{"msg": err})
		return
	}

	c.JSON(200, gin.H{
		"game": saved,
	})
}
