package request

import (
	"github.com/gin-gonic/gin"
	"yyds-pro/model"
	"yyds-pro/service"
	"yyds-pro/service/serviceimpl"
)

type apkController struct {
	service service.ApkSer
}

//
//  NewUserController
//  @Description: apk处理器
//  @param g
//
func NewApkController(g *model.Routes) {
	handler := apkController{service: serviceimpl.Apks}
	g.Public.POST("getApkById", handler.GetApkById)
}

func (a apkController) GetApkById(context *gin.Context) {
	id := 1
	a.service.GetApkById(id)
}
