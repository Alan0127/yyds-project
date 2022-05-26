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
	GetAllApps(ctx *trace.Trace, req model.GetAppsReq) ([]model.AppInfos, error)

	FindApkById(ctx *trace.Trace, id int) (model.AppInfo, error)

	ChangeTaskOrderStatusByOrderInfo(ctx *trace.Trace, orderReq model.OrderReq, cal int) (err error)
}
