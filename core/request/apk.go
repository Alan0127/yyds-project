package request

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"yyds-pro/core"
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
//  @Description: apk处理器
//  @param g
//
func NewApkController(g *model.Routes) {
	handler := apkController{service: serviceimpl.NewApkService()}
	g.Public.POST("getApkById", handler.GetApkById)
}

func (a apkController) GetApkById(c *gin.Context) {
	_, traceCtx := core.GetTrace(c)
	var i int
	err := c.ShouldBindWith(&i, binding.JSON)
	if err != nil {
		log.GetLogger().InfoCtx(traceCtx, err)
	}

	id := 1
	a.service.GetApkById(traceCtx, id)
}
