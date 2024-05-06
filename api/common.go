package api

import (
	"github.com/gin-gonic/gin"
	"internal_api/model"
)

// Dict 字典
// @Summary 字典
// @Tags 公共接口
// @Description 字典
// @Accept application/json
// @Produce application/json
// @Router /api/v1/common/dict [get]
// @Success 200 {object} map[string][]model.CommonDict
// @Security ApiKeyAuth
func Dict(ctx *gin.Context) {
	var dict = make(map[string][]model.CommonDict, 0)
	// 任务状态
	dict["status"] = []model.CommonDict{
		{
			Id:     "1",
			Name:   "运行中",
			Status: 0,
		},
		{
			Id:     "2",
			Name:   "已结束",
			Status: 0,
		},
	}

	// 任务运行结果
	dict["result"] = []model.CommonDict{
		{
			Id:     "1",
			Name:   "无",
			Status: 0,
		},
		{
			Id:     "2",
			Name:   "运行成功",
			Status: 0,
		}, {
			Id:     "3",
			Name:   "运行失败",
			Status: 0,
		}, {
			Id:     "4",
			Name:   "手动取消",
			Status: 0,
		}, {
			Id:     "5",
			Name:   "异常",
			Status: 0,
		},
	}

	// agent status 状态
	dict["agentStatus"] = []model.CommonDict{
		{
			Id:     "200",
			Name:   "running",
			Status: 0,
		}, {
			Id:     "-1",
			Name:   "unknown",
			Status: 0,
		},
	}

	model.Success(ctx, dict, "")
}
