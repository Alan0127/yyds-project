package monitor

import "yyds-pro/log"

//todo 待实现的监控日志功能
func SetMonitorLog() {
	log.InitMonitorLog(log.SetLevel("info"), log.SetPath("/logs/monitor/"))
}
