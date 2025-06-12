package main

import (
	"fmt"
	"go.uber.org/zap"
	"internal_api/global"
	"internal_api/initialize"
	"internal_api/pkg/job_queue/task"
	"log"
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

	// 初始化生产者和消费者
	const redisAddr = "127.0.0.1:6379"
	task.InitProducer(redisAddr)
	task.InitConsumer(redisAddr)
	// 启动消费者
	go func() {
		if err := task.StartWorker(); err != nil {
			log.Fatalf("Worker error: %v", err)
		}
	}()

	// ws socket init
	//go app.InitSocket()

	zap.S().Info(fmt.Sprintf("\naddr: http://localhost:%d \nswagger: http://localhost:%d/swagger/index.html", global.CONFIG.App.Port, global.CONFIG.App.Port))
	//// 启动 web 服务
	//if err := Router.Run(fmt.Sprintf(":%d", global.CONFIG.App.Port)); err != nil {
	//	panic(fmt.Sprintf("启动失败:%s", err.Error()))
	//}

	if err := Router.Run(fmt.Sprintf(":%d", 5001)); err != nil {
		panic(fmt.Sprintf("启动失败:%s", err.Error()))
	}

}
