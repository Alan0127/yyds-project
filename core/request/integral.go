package request

import (
	"errors"
	"github.com/gin-gonic/gin"
	"yyds-pro/core"
	_const "yyds-pro/core/const"
	"yyds-pro/core/response"
	"yyds-pro/model"
	"yyds-pro/service"
	"yyds-pro/service/serviceimpl"
)

type rateController struct {
	service service.RateLimiterService
}

func NewRateController(g *model.Routes) {
	handler := rateController{serviceimpl.NewRateLimiterService()}
	g.Public.POST("initIntegralDb", handler.InitIntegralDb)
	g.Public.POST("getIntegral", handler.GetIntegral)
}

//
//  InitIntegralDb
//  @Description: 初始化integral-db
//  @receiver r
//  @param c
//
func (r rateController) InitIntegralDb(c *gin.Context) {
	_, traceCtx := core.GetTrace(c)
	var integral model.Integral
	err := core.BindReqWithContext(traceCtx, c, &integral)
	if err != nil {
		response.ResError(c, traceCtx, err)
		return
	}
	err = r.service.InitIntegralDb(traceCtx, integral)
	if err != nil {
		response.ResError(c, traceCtx, err)
		return
	}
	response.ResSuccess(c, traceCtx, _const.ResponseSuccess, _const.InitIntegralDbMsg)
}

//
//  GetIntegral
//  @Description: 模拟抢积分活动会出现高并发
//  @receiver r
//  @param c
//
func (r rateController) GetIntegral(c *gin.Context) {
	_, traceCtx := core.GetTrace(c)
	var user model.UserInfo
	err := core.BindReqWithContext(traceCtx, c, &user)
	if err != nil {
		response.ResError(c, traceCtx, err)
		return
	}
	//用户校验相关(暂未处理)
	if user.UserName == "" || user.Password == "" || user.Token == "" || user.Country == "" {
		response.ResError(c, traceCtx, errors.New("account check error"))
		return
	}
	//msg:查询cache返回的描述信息；val:抢到的积分，默认为0； err:错误信息
	msg, val, err := r.service.GetIntegralByUserInfo(traceCtx, user)
	if val != 0 {
		err = r.service.UpdateUserIntegral(traceCtx, val, user.UserName)
		if err != nil {
			response.ResError(c, traceCtx, err)
			return
		}
	}
	if err != nil {
		response.ResError(c, traceCtx, err)
		return
	}
	response.ResSuccess(c, traceCtx, val, msg)
}
