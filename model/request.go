package model

import (
	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	Id          uint
	Name        string
	Email       string
	AuthorityId uint
	jwt.StandardClaims
}

// 操作实体
type OperateCallBackReq struct {
	Code int    `form:"code" json:"code"` // 1-运行 2-kill
	Msg  string `form:"msg" json:"msg"`   // 详细信息
}
