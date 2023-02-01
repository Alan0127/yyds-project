package main

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"yyds-pro/core"
	"yyds-pro/log"
	"yyds-pro/router"
)

func main() {
	core.InitDefault()
	g := router.Init()
	err := g.Run(":8888")
	if err != nil {
		log.GetLogger().Error("start error!", zap.Any("err", err.Error()))
		panic(err)
	}
	//程序终止
	quite()

}

// 优雅退出
func quite() {
	//接收退出信号的chanel
	sig := make(chan os.Signal)
	//指定哪些信号可以转发到chanel，如果没有列出，会将所有信号传递到chanel
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)
	fmt.Printf("接收到的信号: %v \n", <-sig)
	fmt.Println("主goroutine结束")
}
