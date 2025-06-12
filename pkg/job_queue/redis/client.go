package redis

import "github.com/hibiken/asynq"

var RedisClientOpt = asynq.RedisClientOpt{
	Addr: "127.0.0.1:6379",
}
