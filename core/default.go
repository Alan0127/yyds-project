package core

import (
	"go.uber.org/zap"
	"yyds-pro/config"
	"yyds-pro/log"
	"yyds-pro/server/mysql"
	"yyds-pro/server/redis"
)

func InitDefaultConnections() {
	conf, err := config.LoadConfig()
	log.InitDefaultLog(log.SetLevel("info"), log.SetPath("/logs/"))
	l := log.GetLogger()
	if err != nil {
		l.Info("load config error", zap.Any("error", err.Error()))
	}
	sqlErr := mysql.InitMysql(conf) //初始化mysql
	if sqlErr != nil {
		l.Info("connect mysql error!", zap.Any("error", sqlErr.Error()))
	}
	redisErr := redis.InitRedis(conf) //初始化redis
	if redisErr != nil {
		l.Info("connect redis error!", zap.Any("error", redisErr.Error()))
	}
}
