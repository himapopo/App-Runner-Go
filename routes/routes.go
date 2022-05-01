package routes

import "github.com/gin-gonic/gin"

func RouterDifinition(e *gin.Engine, birdRouter BirdRouter) {
	birdRouter.BirdRouterDifinition(e)
}
