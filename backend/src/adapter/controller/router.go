package controller

import (
	"github.com/gin-gonic/gin"
)

type Routing struct {
	Gin *gin.Engine
}

func NewRouting(gameController GameController) *Routing {
	r := &Routing{
		Gin: gin.Default(),
	}
	r.Gin.GET("/ping", func(c *gin.Context) { r.getPing(c) })
	r.Gin.GET("/game/:id", func(c *gin.Context) { gameController.Get(c) })
	r.Gin.POST("/game", func(c *gin.Context) { gameController.Post(c) })
	return r
}

func (r *Routing) getPing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
