package request

import (
	"github.com/gin-gonic/gin"
	"yyds-pro/core"
	"yyds-pro/core/const"
	"yyds-pro/core/response"
	"yyds-pro/model"
	"yyds-pro/service"
	"yyds-pro/service/serviceimpl"
)

type apkController struct {
	service service.ApkService
}

//
//  NewUserController
//  @Description: app处理器
//  @param g
//
func NewAppController(g *model.Routes) {
	handler := apkController{service: serviceimpl.NewApkService()}
	g.Public.POST("getAppAppInfos", handler.GetAllApps)
	g.Public.POST("getAppInfoById", handler.GetApkById)
	g.Public.POST("order", handler.ChangeOrderStatus)
	g.Public.POST("RushPurchase", handler.RushPurchase)
}

//
//  GetAllApps
//  @Description: 获取所有
//  @receiver a
//  @param c
//
func (a apkController) GetAllApps(c *gin.Context) {
	_, traceCtx := core.GetTrace(c)
	var appReq model.GetAppsReq
	err := core.BindReqWithContext(traceCtx, c, &appReq)
	if err != nil {
		response.ResError(c, traceCtx, err)
		return
	}
	res, err := a.service.GetAllApps(traceCtx, appReq)
	if err != nil {
		response.ResError(c, traceCtx, err)
		return
	}
	response.ResSuccess(c, traceCtx, res, _const.GetAllAppsMsg)
}

//
//  GetApkById
//  @Description: 根据appId获取应用详情
//  @receiver a
//  @param c
//
func (a apkController) GetApkById(c *gin.Context) {
	_, traceCtx := core.GetTrace(c)
	var id model.ReqId
	err := core.BindReqWithContext(traceCtx, c, &id)
	if err != nil {
		response.ResError(c, traceCtx, err)
		return
	}
	res, err := a.service.GetApkById(traceCtx, id.Id)
	if err != nil {
		response.ResError(c, traceCtx, err)
		return
	}
	response.ResSuccess(c, traceCtx, res, _const.GetApkByIdMsg)
}

//
//  ChangeOrderStatus
//  @Description: 修改预约状态
//  @receiver a
//  @param c
//
func (a apkController) ChangeOrderStatus(c *gin.Context) {
	_, traceCtx := core.GetTrace(c)
	var orderReq model.OrderReq
	err := core.BindReqWithContext(traceCtx, c, &orderReq)
	if err != nil {
		response.ResError(c, traceCtx, err)
		return
	}
	err = a.service.ChangeTaskOrderStatusByOrderInfo(traceCtx, orderReq)
	if err != nil {
		response.ResError(c, traceCtx, err)
		return
	}
	response.ResSuccess(c, traceCtx, "", _const.ChangeOrderStatusMsg)
}

func (a apkController) RushPurchase(c *gin.Context) {
	_, traceCtx := core.GetTrace(c)
	var userInfo model.UserPurchase
	err := core.BindReqWithContext(traceCtx, c, &userInfo)
	if err != nil {
		response.ResError(c, traceCtx, err)
		return
	}

}
