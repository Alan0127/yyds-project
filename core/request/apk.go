package request

import (
	"github.com/gin-gonic/gin"
	"yyds-pro/core"
	"yyds-pro/core/response"
	"yyds-pro/log"
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
	g.Public.POST("getAppInfoById", handler.GetApkById)
}

func (a apkController) GetApkById(c *gin.Context) {
	_, traceCtx := core.GetTrace(c)
	var id model.ReqId
	err := core.BindReqWithContext(traceCtx, c, &id)
	if err != nil {
		log.L.ErrorCtx(traceCtx, err)
		response.ResError(c, traceCtx, err)
		return
	}
	res, err := a.service.GetApkById(traceCtx, id.Id)
	if err != nil {
		log.L.ErrorCtx(traceCtx, err)
		response.ResError(c, traceCtx, err)
		return
	}
	response.ResSuccess(c, traceCtx, res)
}
