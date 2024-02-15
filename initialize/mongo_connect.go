package initialize

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"internal_api/global"
)

func MongoClient() {
	zap.S().Info("初始化Mongo...")
	// 设置 MongoDB 连接选项
	// mongodb://username:password@host:port/database
	uri := "mongodb://" + global.CONFIG.Mongo.Username + ":" + global.CONFIG.Mongo.Password + "@" + global.CONFIG.Mongo.Host + ":" + global.CONFIG.Mongo.Port

	// 创建 MongoDB 客户端
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	global.MONGO_CLIENT = client
}
