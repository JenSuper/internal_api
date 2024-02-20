package model

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
*
全局返回结果
*/
func buildH(code int, msg string, data any) gin.H {
	return gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	}
}

// Success 成功返回
func Success(ctx *gin.Context, data any, msg string) {
	ctx.JSON(http.StatusOK, buildH(200, msg, data))
}

// Error 成功返回
func Error(ctx *gin.Context, data any, msg string) {
	ctx.JSON(http.StatusOK, buildH(-1, msg, data))
}

func CustomUnauthorized(ctx *gin.Context, dataCode int, data any, msg string) {
	ctx.JSON(http.StatusUnauthorized, buildH(dataCode, msg, data))
}

func Custom(ctx *gin.Context, httpCode int, dataCode int, data any, msg string) {
	ctx.JSON(httpCode, buildH(dataCode, msg, data))
}

// 数据build条件，下载和查询数据使用
type CategoryIds struct {
	C1Ids []string
	C2Ids []string
	C3Ids []string
}

// 数据build条件，下载和查询数据使用
type SystemInfo struct {
	data any
}
