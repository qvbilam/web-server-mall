package initialize

import (
	"github.com/gin-gonic/gin"
	"mall/middleware"
	mallRouter "mall/router"
	"net/http"
)

func InitRouters() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors())
	router.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})

	apiRouter := router.Group("")
	mallRouter.InitMallRouter(apiRouter)

	return router
}
