package usecase

import (
	"App-Runner-go/infrastructure/entity"

	"github.com/gin-gonic/gin"
)

type birdUseCase struct {
}

type BirdUseCase interface {
	Fetch(ctx *gin.Context) *entity.Bird
}

func NewBirdUseCase() BirdUseCase {
	return &birdUseCase{}
}

func (r *birdUseCase) Fetch(ctx *gin.Context) *entity.Bird {
	var b entity.Bird
	err := ctx.BindJSON(&b)
	if err != nil {
		panic(err)
	}
	return &b
}
