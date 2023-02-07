package handler

import (
	"github.com/gin-gonic/gin"
	"yyds-pro/core"
	"yyds-pro/core/const"
	"yyds-pro/core/response"
	"yyds-pro/model"
	"yyds-pro/service"
)

type ApkController struct {
	Service service.ApkService
}

//
//  GetAllApps
//  @Description: 获取所有
//  @receiver a
//  @param c
//
func (a ApkController) GetAllApps(c *gin.Context) {
	_, traceCtx := core.GetTrace(c)
	var appReq model.GetAppsReq
	err := core.BindReqWithContext(traceCtx, &appReq)
	if err != nil {
		response.ResError(traceCtx, err)
		return
	}
	res, err := a.Service.GetAllApps(traceCtx, appReq)
	if err != nil {
		response.ResError(traceCtx, err)
		return
	}
	response.ResSuccess(traceCtx, res, _const.GetAllAppsMsg)
}

//
//  GetApkById
//  @Description: 根据appId获取应用详情
//  @receiver a
//  @param c
//
func (a ApkController) GetApkById(c *gin.Context) {
	_, traceCtx := core.GetTrace(c)
	var id model.ReqId
	err := core.BindReqWithContext(traceCtx, &id)
	if err != nil {
		response.ResError(traceCtx, err)
		return
	}
	res, err := a.Service.GetApkById(traceCtx, id.Id)
	if err != nil {
		response.ResError(traceCtx, err)
		return
	}
	response.ResSuccess(traceCtx, res, _const.GetApkByIdMsg)
}

//
//  ChangeOrderStatus
//  @Description: 修改预约状态
//  @receiver a
//  @param c
//
func (a ApkController) ChangeOrderStatus(c *gin.Context) {
	_, traceCtx := core.GetTrace(c)
	var orderReq model.OrderReq
	err := core.BindReqWithContext(traceCtx, &orderReq)
	if err != nil {
		response.ResError(traceCtx, err)
		return
	}
	err = a.Service.ChangeTaskOrderStatusByOrderInfo(traceCtx, orderReq)
	if err != nil {
		response.ResError(traceCtx, err)
		return
	}
	response.ResSuccess(traceCtx, "", _const.ChangeOrderStatusMsg)
}
