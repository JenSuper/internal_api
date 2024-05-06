package router

import (
	"github.com/gin-gonic/gin"
	"internal_api/api"
)

func InitServiceRouter(Router *gin.RouterGroup) {
	{

	}
}

func InitAuthRouter(Router *gin.RouterGroup) {
	{
		Router.GET("/auth/login", api.LoginByCode)
	}
}

func InitCommonRouter(Router *gin.RouterGroup) {
	{
		Router.GET("/common/dict", api.Dict)
	}
}
