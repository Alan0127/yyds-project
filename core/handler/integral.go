package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"yyds-pro/core"
	_const "yyds-pro/core/const"
	"yyds-pro/core/response"
	"yyds-pro/model"
	"yyds-pro/service"
)

type RateController struct {
	Service service.RateLimiterService
}

//
//  InitIntegralDb
//  @Description: 初始化integral-db
//  @receiver r
//  @param c
//
func (r RateController) InitIntegralDb(c *gin.Context) {
	_, traceCtx := core.GetTrace(c)
	var integral model.Integral
	err := core.BindReqWithContext(traceCtx, &integral)
	if err != nil {
		response.ResError(traceCtx, err)
		return
	}
	err = r.Service.InitIntegralDb(traceCtx, integral)
	if err != nil {
		response.ResError(traceCtx, err)
		return
	}
	response.ResSuccess(traceCtx, _const.ResponseSuccess, _const.InitIntegralDbMsg)
}

//
//  GetIntegral
//  @Description: 模拟抢积分活动会出现高并发
//  @receiver r
//  @param c
//
func (r RateController) GetIntegral(c *gin.Context) {
	_, traceCtx := core.GetTrace(c)
	var user model.UserInfo
	err := core.BindReqWithContext(traceCtx, &user)
	if err != nil {
		response.ResError(traceCtx, err)
		return
	}
	//用户校验相关(暂未处理)
	if user.UserName == "" || user.Password == "" || user.Country == "" {
		response.ResError(traceCtx, errors.New("account check error"))
		return
	}
	//msg:查询cache返回的描述信息；val:抢到的积分，默认为0； err:错误信息
	msg, val, err := r.Service.GetIntegralByUserInfo(traceCtx, user)
	if val != 0 {
		err = r.Service.UpdateUserIntegral(traceCtx, val, user.UserName)
		if err != nil {
			response.ResError(traceCtx, err)
			return
		}
	}
	if err != nil {
		response.ResError(traceCtx, err)
		return
	}
	response.ResSuccess(traceCtx, val, msg)
}
