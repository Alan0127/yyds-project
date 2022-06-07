package core

import (
	"yyds-pro/config"
	"yyds-pro/log"
	"yyds-pro/server/mysql"
	"yyds-pro/server/redis"
)

func InitDefaultConnections() {
	log.InitDefaultLog(log.SetLevel("info"), log.SetPath("/logs/"))
	conf, err := config.LoadConfig()
	if err != nil {
		panic("load config error")
	}
	sqlErr := mysql.InitMysql(conf) //初始化mysql
	if sqlErr != nil {
		panic("connect mysql error!")
	}
	redisErr := redis.InitRedis(conf) //初始化redis
	if redisErr != nil {
		panic("connect redis error!")
	}
}
