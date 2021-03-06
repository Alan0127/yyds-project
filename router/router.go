package router

import (
	"github.com/gin-gonic/gin"
	"yyds-pro/core/request"
	"yyds-pro/middleware"
	"yyds-pro/model"
)

func Init(g *gin.Engine) {
	groupV1 := g.Group("/proApi/v1/")
	request.NewAppController(CreateRoute(groupV1, "/apk/"))
	request.NewRateController(CreateRoute(groupV1, "/rateLimiterAndIntegral/"))
	request.NewUserController(CreateRoute(groupV1, "/user/"))
}

func CreateRoute(group *gin.RouterGroup, path string) *model.Routes {
	tg := group.Group(path)
	tg.Use(middleware.Logger()) //middleware.JwtToken())
	return &model.Routes{
		Public: tg,
	}
}
