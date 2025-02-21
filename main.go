package main

import (
	"fmt"
	"go.uber.org/zap"
	"internal_api/global"
	"internal_api/initialize"
)

// @title Internal API Server
// @version 1.0
// @description Internal API Server 服务
// @termsOfService http://swagger.io/terms/

// @contact.name Admin
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	// 初始化 api
	Router := initialize.Routers()

	//initialize.RedisClient()
	//initialize.MongoClient()

	// ws socket init
	//go app.InitSocket()

	zap.S().Info(fmt.Sprintf("\naddr: http://localhost:%d \nswagger: http://localhost:%d/swagger/index.html", global.CONFIG.App.Port, global.CONFIG.App.Port))
	// 启动 web 服务
	if err := Router.Run(fmt.Sprintf(":%d", global.CONFIG.App.Port)); err != nil {
		panic(fmt.Sprintf("启动失败:%s", err.Error()))
	}

}
