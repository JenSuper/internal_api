package task

const TypePollAPI = "task:poll_api"

type PollAPIPayload struct {
	TaskID string `json:"task_id"`
}
