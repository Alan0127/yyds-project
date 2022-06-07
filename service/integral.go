package service

import (
	"yyds-pro/model"
	"yyds-pro/trace"
)

type RateLimiterService interface {
	//
	//  InitRewardInfoToDb
	//  @Description: 初始化reward service interface
	//  @param id
	//  @param money
	//  @return error
	//
	InitIntegralDb(ctx *trace.Trace, integral model.Integral) error

	//
	//  GetIntegralByUserInfo
	//  @Description: 抢积分逻辑interface
	//  @param ctx
	//  @param userInfo
	//  @return int
	//  @return error
	//
	GetIntegralByUserInfo(ctx *trace.Trace, userInfo model.UserInfo) (string, int, error)

	//
	//  UpdateUserIntegral
	//  @Description: 更新user积分service interface
	//  @param ctx
	//  @param val
	//  @return error
	//
	UpdateUserIntegral(ctx *trace.Trace, val int, userName string) error
}
