package initialize

import (
	"go.uber.org/zap"
	"internal_api/global"
	"internal_api/repository"
)

func RedisClient() {
	zap.S().Infof("初始化Redis..")
	rep, err := repository.RedisConnect()
	if err != nil {
		zap.S().Error("Redis Init Error")
	}
	global.REDIS_CLIENT = rep
}
