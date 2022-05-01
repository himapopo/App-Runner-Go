package routes

import (
	"App-Runner-go/controller"

	"github.com/gin-gonic/gin"
)

type BirdRouter interface {
	BirdRouterDifinition(e *gin.Engine)
}

type birdRouter struct {
	controller controller.BirdController
}

func NewBirdRouter(controller controller.BirdController) BirdRouter {
	return &birdRouter{
		controller: controller,
	}
}

func (r *birdRouter) BirdRouterDifinition(e *gin.Engine) {
	g := e.Group("/bird")
	g.GET("/", r.controller.Fetch())
}
