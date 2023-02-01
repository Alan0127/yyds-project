package service

import (
	"yyds-pro/model/report"
	"yyds-pro/trace"
)

//
//  ReportService
//  @Description: report service接口
//
type ReportService interface {
	//
	//  GenerateReport
	//  @Description: 生成报表并上传到oss
	//  @param c
	//  @param req
	//  @return code
	//  @return err
	//
	GenerateReport(c *trace.Trace, req report.GenerateReport) (code int, err error)

	//
	//  GetReport
	//  @Description: 从oss下载报表
	//  @param c
	//  @return code
	//  @return err
	//
	GetReport(c *trace.Trace) (code int, err error)

	//
	//  GenerateReportByKafka
	//  @Description: 读kafka生成报表
	//  @param c
	//  @return code
	//  @return err
	//
	GenerateReportByKafka(c *trace.Trace, req report.StartFlag) (code int, err error)

	//
	//  SendToElastic
	//  @Description: 发送数据到elastic中
	//  @param c
	//  @param req
	//  @return code
	//  @return err
	//
	SendToElastic(c *trace.Trace, req report.ElasticReq) (code int, err error)
}
