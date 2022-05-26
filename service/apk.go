package service

import (
	"yyds-pro/model"
	"yyds-pro/trace"
)

//
//  ApkService
//  @Description:ApkService抽象接口
//
type ApkService interface {
	//
	//  GetAllApps
	//  @Description: GetAllApps接口方法
	//  @param ctx
	//  @param req
	//  @return model.AppInfos
	//  @return error
	//
	GetAllApps(ctx *trace.Trace, req model.GetAppsReq) ([]model.AppInfos, error)
	//
	//  GetApkById
	//  @Description: GetApkById接口方法
	//  @param ctx
	//  @param id
	//  @return model.AppInfo
	//  @return error
	//
	GetApkById(ctx *trace.Trace, id int) (model.AppInfo, error)

	//
	//  ChangeTaskOrderStatusByOrderInfo
	//  @Description: ChangeTaskOrderStatusByOrderInfo接口方法
	//  @param ctx
	//  @param orderInfo
	//  @return model.OrderRes
	//  @return error
	//
	ChangeTaskOrderStatusByOrderInfo(ctx *trace.Trace, orderInfo model.OrderReq) (err error)
}
