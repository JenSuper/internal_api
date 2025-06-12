package task

import (
	"encoding/json"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

var Client *asynq.Client

func InitProducer(redisAddr string) {
	Client = asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
}

func EnqueuePollTask(taskID string) error {
	payload, err := json.Marshal(PollAPIPayload{TaskID: taskID})
	if err != nil {
		return err
	}

	task := asynq.NewTask(TypePollAPI, payload)

	// 加入队列
	info, err := Client.Enqueue(task, asynq.MaxRetry(10), asynq.Timeout(30*time.Second))
	if err != nil {
		return err
	}
	log.Printf("Enqueued task: id=%s queue=%s", info.ID, info.Queue)
	return nil
}
