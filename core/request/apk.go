package request

import (
	"fmt"
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
//  @Description: app处理器
//  @param g
//
func NewAppController(g *model.Routes) {
	handler := apkController{service: serviceimpl.NewApkService()}
	g.Public.POST("getAppInfoById", handler.GetApkById)
}

func (a apkController) GetApkById(c *gin.Context) {
	_, traceCtx := core.GetTrace(c)
	var i int
	err := c.ShouldBindWith(&i, binding.JSON)
	if err != nil {
		log.GetLogger().InfoCtx(traceCtx)
	}

	id := 1
	res, err := a.service.GetApkById(traceCtx, id)
	fmt.Println(res)
}
