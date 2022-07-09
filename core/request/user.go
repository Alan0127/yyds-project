package request

import (
	"errors"
	"github.com/gin-gonic/gin"
	"yyds-pro/core"
	"yyds-pro/core/response"
	"yyds-pro/core/util"
	"yyds-pro/middleware"
	"yyds-pro/model"
	"yyds-pro/service"
	"yyds-pro/service/serviceimpl"
)

type UserController struct {
	service service.UserService
}

func NewUserController(g *model.Routes) {
	handler := UserController{service: serviceimpl.NewUserService()}
	g.Public.POST("login", handler.Login)
	g.Public.POST("register", handler.Register)
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
	err := core.BindReqWithContext(traceCtx, c, &user)
	if err != nil {
		response.ResError(c, traceCtx, err)
		return
	}
	//查询
	res, err := u.service.CheckLogin(traceCtx, user.UserName, user.Country)
	if err != nil {
		response.ResError(c, traceCtx, err)
		return
	}
	if res.UserName == "" && res.Password == "" && res.Country == "" {
		response.ResError(c, traceCtx, errors.New("用户不存在！"))
		return
	}
	//密码校验
	flag := util.ComparePwd(res.Password, []byte(user.Password))
	if !flag {
		response.ResError(c, traceCtx, errors.New("密码错误！"))
		return
	}
	token, code := middleware.SetToken(user.UserName)
	if code != 200 {
		response.ResError(c, traceCtx, errors.New("创建token失败！"))
		return
	}
	response.ResSuccess(c, traceCtx, token, "创建token成功！")
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
	err := core.BindReqWithContext(traceCtx, c, &registerInfo)
	if err != nil {
		response.ResError(c, traceCtx, err)
		return
	}
	//查询用户是否存在
	res, err := u.service.CheckLogin(traceCtx, registerInfo.UserName, registerInfo.Country)
	if err != nil {
		response.ResError(c, traceCtx, err)
		return
	}
	if res.UserName != "" { //已注册
		response.ResError(c, traceCtx, errors.New("用户已注册"))
		return
	}
	//加密
	pwd := util.HashAndSalt([]byte(registerInfo.UserPass))
	insertInfo := registerInfo
	insertInfo.UserPass = pwd
	//插入数据
	v, err := u.service.RegisterUser(traceCtx, insertInfo)
	if err != nil {
		response.ResError(c, traceCtx, err)
		return
	}
	if v != 0 {
		response.ResError(c, traceCtx, errors.New("添加用户失败！"))
		return
	}
	response.ResSuccess(c, traceCtx, v, "添加用户成功！")
}
