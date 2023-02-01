package core

import (
	"yyds-pro/config"
	v "yyds-pro/core/const"
	"yyds-pro/log"
	"yyds-pro/model"
	"yyds-pro/monitor"
	"yyds-pro/server/mysql"
	"yyds-pro/server/oss"
	"yyds-pro/server/redis"
)

var (
	DefaultConfig model.AppConfig
)

func InitDefault() {
	v.InitDefaultLanguage()
	conf, err := config.LoadConfig()
	DefaultConfig = conf
	//初始化日志
	log.InitDefaultLog(log.SetLevel("info"), log.SetPath("/logs/"))
	monitor.SetMonitorLog() //监控日志
	if err != nil {
		panic("load config error")
	}
	//初始化mysql连接
	err = mysql.InitMysql(conf) //初始化mysql
	if err != nil {
		panic("connect mysql error!")
	}
	//初始化redis连接
	err = redis.InitRedis(conf) //初始化redis
	if err != nil {
		panic("connect redis error!")
	}
	//初始化oss
	err = oss.InitOss(conf)
	if err != nil {
		panic("connect oss error!")
	}
	//初始化elastic
	//err = elastic.InitElastic(conf)
	//if err != nil {
	//	panic("connect elastic error!")
	//}
}
