package initialize

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "internal_api/docs"
	"internal_api/global"
	"internal_api/middlewares"
	"internal_api/router"
	"net/http"
)

func Routers() *gin.Engine {
	// 模式选择
	gin.SetMode(chooseMode())

	Router := gin.Default()
	//配置跨域
	Router.Use(middlewares.Cors())

	Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

	// swagger http://localhost:17000/swagger/index.html
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ApiGroup := Router.Group("/api/v1")
	//ApiGroup := Router.Group("/api/v1", middlewares.JWTAuth())
	// 服务初始化
	router.InitAuthRouter(ApiGroup)
	router.InitCommonRouter(ApiGroup)
	router.RedisRouter(ApiGroup)

	return Router
}

// chooseMode 设置 gin 的启动模式
func chooseMode() string {
	switch global.ENV {
	case "local":
		return gin.DebugMode
	case "test":
		return gin.TestMode
	case "prod":
		return gin.ReleaseMode
	default:
		return gin.DebugMode
	}

}
