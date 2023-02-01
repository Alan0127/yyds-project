package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"yyds-pro/core"
	_const "yyds-pro/core/const"
	"yyds-pro/core/response"
	"yyds-pro/model/report"
	"yyds-pro/server/kafka"
	"yyds-pro/service"
)

type ReportController struct {
	Service service.ReportService
}

//
//  TestProducerKafka
//  @Description: kafka测试相关功能接口
//  @receiver r
//  @param c
//
func (r *ReportController) TestProducerKafka(c *gin.Context) {
	_, traceCtx := core.GetTrace(c)
	var testProducer report.TestProducer
	err := core.BindReqWithContext(traceCtx, &testProducer)
	if err != nil {
		response.ResError(traceCtx, err)
		return
	}
	kafkaProducer, err := kafka.CreateProducer(core.DefaultConfig.App.Kafka.Address+":"+core.DefaultConfig.App.Kafka.Port, "test1", 5, "Sync")
	if err != nil {
		response.ResError(traceCtx, err)
		return
	}
	v, _ := json.Marshal(testProducer)
	err = kafkaProducer.Send(v)
	if err != nil {
		response.ResError(traceCtx, err)
		return
	}
	response.ResSuccess(traceCtx, "send kafka-producer successful", _const.SendKafkaSuccessful)
}

//
//  GenerateReport
//  @Description: 生成报表上传oss
//  @receiver r
//  @param c
//
func (r *ReportController) GenerateReport(c *gin.Context) {
	_, traceCtx := core.GetTrace(c)
	var req report.GenerateReport
	err := core.BindReqWithContext(traceCtx, &req)
	if err != nil {
		response.ResError(traceCtx, err)
		return
	}
	_, err = r.Service.GenerateReport(traceCtx, req)
	if err != nil {
		fmt.Println("the err is: ", err)
		response.ResError(traceCtx, err)
		return
	}
	response.ResSuccess(traceCtx, nil, "create excel successful")
}

//
//  GetReport
//  @Description: 从oss下载报告
//  @receiver r
//  @param c
//
func (r *ReportController) GetReport(c *gin.Context) {
	_, traceCtx := core.GetTrace(c)
	_, err := r.Service.GetReport(traceCtx)
	if err != nil {
		response.ResError(traceCtx, err)
	}
	response.ResSuccess(traceCtx, nil, "download file success")
}

//
//  GenerateReportByKafka
//  @Description: 从kafka中读取数据生成文档上传至oss
//  @receiver r
//  @param c
//
func (r *ReportController) GenerateReportByKafka(c *gin.Context) {
	_, traceCtx := core.GetTrace(c)
	var req report.StartFlag
	err := core.BindReqWithContext(traceCtx, &req)
	if err != nil {
		response.ResError(traceCtx, err)
		return
	}
	_, err = r.Service.GenerateReportByKafka(traceCtx, req)
	if err != nil {
		response.ResError(traceCtx, err)
		return
	}
	response.ResSuccess(traceCtx, nil, "generate by kafka success")
}

//
// SendToElastic
//  @Description: elastic数据收集
//  @receiver r
//  @param c
//
func (r *ReportController) SendToElastic(c *gin.Context) {
	_, traceCtx := core.GetTrace(c)
	var req report.ElasticReq
	err := core.BindReqWithContext(traceCtx, &req)
	if err != nil {
		response.ResError(traceCtx, err)
		return
	}
	_, err = r.Service.SendToElastic(traceCtx, req)
	if err != nil {
		return
	}
	response.ResSuccess(traceCtx, nil, "send msg to elastic success")
}
