package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"internal_api/model"
)

// List list
// @Summary Admin
// @Tags Admin
// @Description Admin
// @Accept application/json
// @Produce application/json
// @Router /api/v1/admin/list [get]
// @Success 200 {object} map[string][]model.CommonDict
// @Security ApiKeyAuth
func List(ctx *gin.Context) {
	zap.S().Info("admin todo list")
	model.Success(ctx, nil, "")
}
