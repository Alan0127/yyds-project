package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"yyds-pro/core"
	"yyds-pro/log"
	"yyds-pro/router"
)

func main() {
	core.InitDefaultConnections()
	g := gin.New()
	router.Init(g)
	err := g.Run(":8888")
	if err != nil {
		log.GetLogger().Info("start error!", zap.Any("err", err.Error()))
	}
}
