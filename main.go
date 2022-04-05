package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"yyds-pro/config"
	"yyds-pro/log"
	"yyds-pro/router"
	"yyds-pro/server/mysql"
	"yyds-pro/server/redis"
)

//初始化工作
func init() {
	conf, err := config.LoadConfig()
	log.InitDefaultLog(log.SetLevel("info"), log.SetPath("/logs/"))
	l := log.GetLogger()
	if err != nil {
		l.Info("load config error", zap.Any("error", err.Error()))
	}
	sqlErr := mysql.InitMysql(conf)   //初始化mysql
	redisErr := redis.InitRedis(conf) //初始化redis
	if sqlErr != nil {
		l.Info("connect mysql error!", zap.Any("error", sqlErr.Error()))
	}
	if redisErr != nil {
		l.Info("connect redis error!", zap.Any("error", redisErr.Error()))
	}
}

func main() {
	g := gin.New()
	router.Init(g)
	err := g.Run(":8888")
	if err != nil {
		log.GetLogger().Info("start error!", zap.Any("err", err.Error()))
	}
}
