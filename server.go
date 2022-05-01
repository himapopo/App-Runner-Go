package main

import (
	"App-Runner-go/controller"
	"App-Runner-go/routes"
	"App-Runner-go/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	e.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "server run",
		})
	})
	birdUsecase := usecase.NewBirdUseCase()

	birdController := controller.NewBirdController(birdUsecase)

	birdRouter := routes.NewBirdRouter(birdController)
	routes.RouterDifinition(e, birdRouter)

	e.Run(":8080")
}
