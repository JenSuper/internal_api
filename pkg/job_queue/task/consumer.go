package task

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

var Server *asynq.Server

func InitConsumer(redisAddr string) {
	Server = asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{
			Concurrency: 5,
			Queues: map[string]int{
				"default": 1,
			},
		},
	)
}

func StartWorker() error {
	mux := asynq.NewServeMux()
	mux.HandleFunc(TypePollAPI, handlePollTask)

	return Server.Run(mux)
}

func handlePollTask(ctx context.Context, t *asynq.Task) error {
	var p PollAPIPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	log.Printf("Polling for task %s...", p.TaskID)

	for {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		status, err := mockCheckExternalAPI(p.TaskID)
		if err != nil {
			log.Printf("Error polling: %v", err)
			time.Sleep(2 * time.Second)
			continue
		}

		if status == "done" {
			log.Printf("✅ Task %s completed", p.TaskID)
			return nil
		}

		log.Printf("⏳ Task %s not done yet, retrying...", p.TaskID)
		time.Sleep(3 * time.Second)
	}
}

// 模拟调用外部接口
func mockCheckExternalAPI(taskID string) (string, error) {
	// 你可以换成实际接口
	if time.Now().Unix()%5 == 0 {
		return "done", nil
	}
	return "processing", nil
}
