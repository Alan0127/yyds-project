package repository

import (
	"yyds-pro/model"
	"yyds-pro/trace"
)

//
//  ApkRepo
//  @Description: repo抽象接口
//
type ApkRepo interface {
	FindApkById(ctx *trace.Trace, id int) (model.AppInfo, error)

	ChangeTaskOrderStatusByOrderInfo(ctx *trace.Trace, orderReq model.OrderReq, cal uint) (err error)
}
