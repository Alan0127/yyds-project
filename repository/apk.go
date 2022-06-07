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
	//
	//  GetAllApps
	//  @Description: GetAllApps repository interface
	//  @param ctx
	//  @param req
	//  @return []model.AppInfos
	//  @return error
	//
	GetAllApps(ctx *trace.Trace, req model.GetAppsReq) ([]model.AppInfos, error)

	//
	//  FindApkById
	//  @Description: FindApkById repository interface
	//  @param ctx
	//  @param id
	//  @return model.AppInfo
	//  @return error
	//
	FindApkById(ctx *trace.Trace, id int) (model.AppInfo, error)

	//
	//  ChangeTaskOrderStatusByOrderInfo
	//  @Description: ChangeTaskOrderStatusByOrderInfo repository interface
	//  @param ctx
	//  @param orderReq
	//  @param cal
	//  @return err
	//
	ChangeTaskOrderStatusByOrderInfo(ctx *trace.Trace, orderReq model.OrderReq, cal int) (err error)

	GetTaskUserOrderStatus(ctx *trace.Trace, orderReq model.OrderReq) (status int, err error)
}
