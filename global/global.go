package global

import (
	"database/sql"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	CONFIG       Config
	REDIS_CLIENT *redis.Client
	MONGO_CLIENT *mongo.Client
	MYSQL_CLIENT *sql.DB
	ENV          string // 环境变量
)
