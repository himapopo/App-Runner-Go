package controller

import (
	"App-Runner-go/usecase"

	"github.com/gin-gonic/gin"
)

type birdController struct {
	usecase usecase.BirdUseCase
}

type BirdController interface {
	Fetch() gin.HandlerFunc
}

func NewBirdController(usecase usecase.BirdUseCase) BirdController {
	return &birdController{
		usecase: usecase,
	}
}

func (c *birdController) Fetch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		b := c.usecase.Fetch(ctx)
		ctx.JSON(200, b)
	}
}
