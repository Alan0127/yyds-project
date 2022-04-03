package main

import (
	"fmt"
	"yyds-pro/config"
	"yyds-pro/log"
	"yyds-pro/server/mysql"
	"yyds-pro/server/redis"
)

//初始化工作
func init() {
	conf, err := config.LoadConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
	sqlErr := mysql.InitMysql(conf)
	redisErr := redis.InitRedis(conf)
	if sqlErr != nil || redisErr != nil {
		fmt.Println(sqlErr, redisErr)
	}
	log.InitDefaultLog(log.SetLevel("info"))
}

func main() {
	//g := gin.New()
}
