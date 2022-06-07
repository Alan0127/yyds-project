package repository

import (
	"yyds-pro/trace"
)

type RateLimiterRepository interface {
	//
	//  InitRewardInfoToDb
	//  @Description: 初始化integral信息repository interface
	//  @param id
	//  @param money
	//  @return error
	//
	InitIntegralDb(ctx *trace.Trace, id, money int) error

	//
	//  UpdateUserIntegral
	//  @Description: 更新用户积分 repository interface
	//  @param ctx
	//  @param val
	//  @return error
	//
	UpdateUserIntegral(ctx *trace.Trace, val int, userName string) error

	//GetTaskUserOrderStatus(ctx *trace.Trace, orderReq model.OrderReq) (int, error)
}
