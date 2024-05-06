package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"internal_api/global"
	"internal_api/middlewares"
	"internal_api/model"
	"internal_api/repository"
	"internal_api/utils"
	"strconv"
	"time"
)

// LoginByCode 通过code登录
// @Summary 登录接口
// @Tags 权限相关接口
// @Description 通过code登录，有效期为5分钟
// @Accept application/json
// @Produce application/json
// @Router /api/v1/auth/login [get]
// @Success 200 {object} string
func LoginByCode(c *gin.Context) {
	// code
	code := c.Query(global.Code)
	//code := c.Request.Header.Get(global.Code)
	deCode, err := utils.AesDecrypt(code)
	if err != nil {
		zap.S().Info(err)
		model.CustomUnauthorized(c, -1, nil, "非法请求")
		c.Abort()
		return
	}

	// 读取json
	var codeInfo = model.CodeInfo{}

	err = json.Unmarshal([]byte(deCode), &codeInfo)
	if err != nil || codeInfo.Email == "" {
		model.CustomUnauthorized(c, -1, nil, "请重新登录")
		c.Abort()
		return
	}

	// 判断是否过期 5 分钟过期
	if time.Now().UnixMilli() > int64(codeInfo.Timestamp) {
		model.CustomUnauthorized(c, -1, nil, "授权失效,请重新登录")
		c.Abort()
		return
	}

	// 生成token
	useId := codeInfo.Id
	j := middlewares.NewJWT()
	token := j.BuildClaims(useId, codeInfo.Name, codeInfo.Email)

	r := repository.New(global.REDIS_CLIENT)

	tokenKey := global.BuildMultiKeys(global.Token, strconv.Itoa(useId))
	err = r.Set(tokenKey, token, global.LoginExpireTime)
	if err != nil {
		zap.S().Error("redis set error")
		return
	}
	codeInfo.Token = token

	model.Success(c, codeInfo, "")

}
