package handler

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"time"
	"yyds-pro/core"
	a "yyds-pro/core/const"
	"yyds-pro/core/response"
	"yyds-pro/core/util"
	"yyds-pro/middleware"
	"yyds-pro/model"
	"yyds-pro/server/redis"
	"yyds-pro/service"
)

type UserController struct {
	Service service.UserService
}

//
//  Login
//  @Description: 用户登录
//  @receiver u
//  @param c
//
func (u UserController) Login(c *gin.Context) {
	_, traceCtx := core.GetTrace(c)
	var user model.UserInfo
	var sessionUserInfo middleware.SessionInfo //cache缓存登录信息，30分钟过期
	err := core.BindReqWithContext(traceCtx, &user)
	if err != nil {
		response.ResError(traceCtx, err)
		return
	}
	//查询
	res, err := u.Service.CheckLogin(traceCtx, user.UserName, user.Country)
	if err != nil {
		response.ResError(traceCtx, err)
		return
	}
	if res.UserName == "" && res.Password == "" && res.Country == "" {
		response.ResError(traceCtx, errors.New("用户不存在！"))
		return
	}
	//密码校验
	flag := util.ComparePwd(res.Password, []byte(user.Password))
	if !flag {
		response.ResError(traceCtx, errors.New("密码错误！"))
		return
	}
	token, code := middleware.SetToken(user.UserName)
	if code != 200 {
		response.ResError(traceCtx, errors.New("创建token失败！"))
		return
	}
	sessionUserInfo.UserName = user.UserName
	sessionUserInfo.UserID = "1"
	temp, _ := json.Marshal(sessionUserInfo)
	err = redis.DefaultRedisClient.Set(a.RedisLoginToken+token, temp, 30*time.Minute).Err()
	if err != nil {
		response.ResError(traceCtx, err)
		return
	}
	response.ResSuccess(traceCtx, token, "创建token成功！")
}

//
//  Register
//  @Description: 用户注册
//  @receiver u
//  @param c
//
func (u UserController) Register(c *gin.Context) {
	_, traceCtx := core.GetTrace(c)
	var registerInfo model.RegisterInfo
	err := core.BindReqWithContext(traceCtx, &registerInfo)
	if err != nil {
		response.ResError(traceCtx, err)
		return
	}
	//查询用户是否存在
	res, err := u.Service.CheckLogin(traceCtx, registerInfo.UserName, registerInfo.Country)
	if err != nil {
		response.ResError(traceCtx, err)
		return
	}
	if res.UserName != "" { //已注册
		response.ResError(traceCtx, errors.New("用户已注册"))
		return
	}
	//加密
	pwd := util.HashAndSalt([]byte(registerInfo.UserPass))
	insertInfo := registerInfo
	insertInfo.UserPass = pwd
	//插入数据
	v, err := u.Service.RegisterUser(traceCtx, insertInfo)
	if err != nil {
		response.ResError(traceCtx, err)
		return
	}
	if v != 0 {
		response.ResError(traceCtx, errors.New("添加用户失败！"))
		return
	}
	response.ResSuccess(traceCtx, v, "添加用户成功！")
}
