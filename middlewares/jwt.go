package middlewares

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"internal_api/global"
	"internal_api/model"
	"internal_api/repository"
	"strconv"
	"strings"
	"time"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 白名单
		if checkWhiteUrl(c.Request.RequestURI) {
			// 继续执行下一个 handle
			c.Next()
			return
		}

		// 获取token
		token := c.Request.Header.Get(global.Authorization)
		zap.S().Info(token)
		if token == "" {
			model.CustomUnauthorized(c, -1, nil, "未登录")
			// 结束上下文，不执行下一个 handle
			c.Abort()
			return
		}

		// parseToken 解析token包含的信息
		j := NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil || claims == nil {
			model.CustomUnauthorized(c, -1, nil, "未登录")
			c.Abort()
		}

		// 根据用户id，取Redis的token
		rc := repository.New(global.REDIS_CLIENT)
		tokenKey := global.BuildMultiKeys(global.Token, strconv.Itoa(int(claims.Id)))
		redisToken, err := rc.Get(tokenKey)
		if err != nil || redisToken == "" {
			model.CustomUnauthorized(c, -1, nil, "登录过期")
			c.Abort()
			return
		}

		// 比较redis token  和 参数token
		if token != redisToken {
			model.CustomUnauthorized(c, -1, nil, "登录过期")
			c.Abort()
			return
		}

		// 更新Redis
		err = rc.Set(tokenKey, token, global.LoginExpireTime)

		c.Set(global.Claims, claims)
		c.Set("userId", claims.Id)
		c.Next()
	}
}

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.CONFIG.Jwt.SigningKey), //可以设置过期时间
	}
}

// 构建claim,生成token
func (j *JWT) BuildClaims(id int, name, email string) string {
	claims := model.CustomClaims{
		Id:    uint(id),
		Name:  name,
		Email: email,
		//AuthorityId:    uint(rsp.Role),
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),               //签名的生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*30, //30天过期
			Issuer:    email,
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		zap.S().Error("token 生成失败")
		return ""
	}
	return token
}

// 创建一个token
func (j *JWT) CreateToken(claims model.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (*model.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*model.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid

	}

}

// 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &model.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*model.CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}

/*
*
验证白名单
*/
func checkWhiteUrl(url string) bool {
	if url == "" {
		return false
	}
	// 判断
	whiteList := strings.Split(global.CONFIG.Jwt.WhiteList, ",")
	for _, s := range whiteList {
		if s != "" && strings.HasPrefix(url, s) {
			return true
		}
	}
	return false
}

/*
*
获取当前用户
*/
func GetClaims(ctx *gin.Context) string {
	cla, _ := ctx.Get(global.Claims)
	if cla == nil {
		return ""
	}
	return cla.(*model.CustomClaims).Name
}

/*
*
获取当前用户
*/
func GetCustomClaims(ctx *gin.Context) *model.CustomClaims {
	cla, _ := ctx.Get(global.Claims)
	if cla == nil {
		return &model.CustomClaims{}
	}
	return cla.(*model.CustomClaims)
}
