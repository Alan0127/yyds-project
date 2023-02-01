package middleware

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	a "yyds-pro/core/const"
	"yyds-pro/core/response"
	"yyds-pro/log"
	"yyds-pro/model"
	"yyds-pro/server/redis"
	"yyds-pro/trace"
)

type SessionInfo struct {
	UserID   int64  `json:"userId"`   // 用户ID
	UserName string `json:"userName"` // 用户名
}

//
// CheckLogin
//  @Description: 登录状态检查
//  @param c
//  @return info
//  @return err
//
func CheckLogin(c *gin.Context) (info SessionInfo, err error) {
	var (
		userInfo  model.UserResInfo
		cacheInfo string
	)
	token := c.GetHeader(a.HeaderToken)
	if token == "" {
		err = errors.New("AuthorizationError")
		return
	}
	cacheInfo, err = redis.DefaultRedisClient.Get(a.RedisLoginToken + token).Result()
	if err != nil {
		return
	}
	if cacheInfo == "" {
		err = errors.New("无缓存用户信息，请检查")
		return
	}
	err = json.Unmarshal([]byte(cacheInfo), &userInfo)
	info.UserName = userInfo.UserName
	info.UserID = userInfo.Id
	if err != nil {
		return
	}
	return
}

//
// WrapHandler
//  @Description: WrapHandler
//  @param handler
//  @return gin.HandlerFunc
//
func WrapHandler(handler func(*gin.Context) (info SessionInfo, err error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		userInfo, err := handler(c)
		if err != nil {
			log.L.ErrorCtx(trace.NewTraceContext(c), err)
			response.ResError(trace.NewTraceContext(c), errors.New("请重新登录"))
		}
		c.Set(a.SessionUserInfo, userInfo)
	}
}
