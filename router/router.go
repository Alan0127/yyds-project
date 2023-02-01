package router

import (
	"github.com/gin-gonic/gin"
	"yyds-pro/core/handler"
	"yyds-pro/middleware"
	"yyds-pro/model"
	"yyds-pro/service/serviceimpl"
)

func Init() *gin.Engine {
	gin.SetMode("debug")
	g := gin.New()
	groupV1 := g.Group("/proApi/v1/")
	NewAppController(CreateRoute(groupV1, "/apk/", middleware.WrapHandler(middleware.CheckLogin))) //登录状态检查
	NewRateController(CreateRoute(groupV1, "/rateLimiterAndIntegral/"))
	NewUserController(CreateRoute(groupV1, "/user/"))
	NewReportController(CreateRoute(groupV1, "/report/"))
	return g
}

func CreateRoute(group *gin.RouterGroup, path string, handlerFunc ...gin.HandlerFunc) *model.Routes {
	tg := group.Group(path, handlerFunc...)
	tg.Use(middleware.Logger())
	return &model.Routes{
		Public: tg,
	}
}

// app模块处理器
func NewAppController(g *model.Routes) {
	h := handler.ApkController{Service: serviceimpl.NewApkService()}
	g.Public.POST("getAppAppInfos", h.GetAllApps)
	g.Public.POST("getAppInfoById", h.GetApkById)
	g.Public.POST("order", h.ChangeOrderStatus)
	//g.Public.POST("RushPurchase", handler.RushPurchase)
}

//福利模块处理器
func NewRateController(g *model.Routes) {
	h := handler.RateController{Service: serviceimpl.NewRateLimiterService()}
	g.Public.POST("initIntegralDb", h.InitIntegralDb)
	g.Public.POST("getIntegral", h.GetIntegral)
}

//报表模块处理器
func NewReportController(g *model.Routes) {
	h := handler.ReportController{Service: serviceimpl.NewReportService()}
	g.Public.POST("testKafka", h.TestProducerKafka)
	g.Public.POST("generateReport", h.GenerateReport)
	g.Public.GET("getReport", h.GetReport)
	g.Public.POST("generateKafkaReport", h.GenerateReportByKafka)
	g.Public.POST("sendToElastic", h.SendToElastic)
}

//用户模块处理器
func NewUserController(g *model.Routes) {
	h := handler.UserController{Service: serviceimpl.NewUserService()}
	g.Public.POST("login", h.Login)
	g.Public.POST("register", h.Register)
}
