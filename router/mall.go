package router

import (
	"github.com/gin-gonic/gin"
	"mall/api"
	"mall/middleware"
)

func InitMallRouter(Router *gin.RouterGroup) {
	MessageRouter := Router.Group("mall/").Use(middleware.Cors()).Use(middleware.Auth())
	{
		MessageRouter.GET("category", api.Category)
		MessageRouter.GET("goods", api.GoodsList)
		MessageRouter.GET("goods/:id", api.GoodsDetail)
	}
}
