package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"internal_api/model"
	"internal_api/pkg/job_queue/task"
	"net/http"
)

// Push push
// @Summary redis push
// @Tags Redis
// @Description Redis
// @Accept application/json
// @Produce application/json
// @Router /api/v1/redis/push [get]
// @Success 200 {object}
func Push(ctx *gin.Context) {
	r := ctx.Request
	w := ctx.Writer

	// @Security ApiKeyAuth
	zap.S().Info("admin todo list")
	taskID := r.URL.Query().Get("id")
	if taskID == "" {
		http.Error(w, "Missing id param", http.StatusBadRequest)
		return
	}
	err := task.EnqueuePollTask(taskID)
	if err != nil {
		http.Error(w, "Failed to enqueue task", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Task submitted.\n"))

	model.Success(ctx, nil, "")
}
