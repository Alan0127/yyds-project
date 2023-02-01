package serviceimpl

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	oss2 "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/xuri/excelize/v2"
	"os"
	"strconv"
	"sync"
	"time"
	"yyds-pro/core"
	_const "yyds-pro/core/const"
	"yyds-pro/core/util"
	"yyds-pro/log"
	"yyds-pro/model/report"
	"yyds-pro/repository"
	"yyds-pro/repository/repoimpl"
	"yyds-pro/server/elastic"
	"yyds-pro/server/kafka"
	"yyds-pro/server/oss"
	"yyds-pro/trace"
)

type ReportService struct {
	ReportRepo repository.ReportRepo
}

func NewReportService() *ReportService {
	return &ReportService{
		ReportRepo: repoimpl.NewReportRepository(),
	}
}

//
//  GenerateReport
//  @Description: 生成报表文件并上传至oss
//  @receiver r
//  @param ctx
//  @param req
//  @return code
//  @return err
//
func (r *ReportService) GenerateReport(ctx *trace.Trace, req report.GenerateReport) (code int, err error) {
	var (
		ch   = make(chan struct{}, 1)
		body []byte
	)
	defer close(ch)
	data := report.ResReportData{}
	c, cancelFunc := context.WithTimeout(ctx.Context, 5*time.Second)
	defer cancelFunc()
	//http调用
	body, err = util.DoHttp(ch, _const.Post, _const.UrlPost, req)
	if err != nil {
		return
	}
	go func() {
		for {
			select {
			case <-ch:
				err = json.Unmarshal(body, &data)
				if err != nil {
					return
				}
				if data.Data == nil {
					err = errors.New("返回数据为空，请检查！")
					return
				}
				//生成报表
				file := excelize.NewFile()
				//删除默认sheet
				defer file.Close()
				file.NewSheet("test_excel")
				file.DeleteSheet("Sheet1")
				rowsCount := len(data.Data)
				for i := 1; i <= rowsCount; i++ {
					_ = file.SetCellValue("test_excel", fmt.Sprintf("A%d", i), strconv.FormatInt(data.Data[i-1].Id, 10))
					_ = file.SetCellValue("test_excel", fmt.Sprintf("B%d", i), data.Data[i-1].PhoneNumber)
					_ = file.SetCellValue("test_excel", fmt.Sprintf("C%d", i), data.Data[i-1].Name)
					_ = file.SetCellValue("test_excel", fmt.Sprintf("D%d", i), data.Data[i-1].Blog)
				}
				bucketName := core.DefaultConfig.App.Oss.BucketName
				bucket, err1 := oss.DefaultOssClient.Bucket(bucketName)
				if err1 != nil {
					err = err1
					return
				}
				//上传至oss
				err = file.SaveAs("./docs/test.xlsx")
				if err != nil {
					return
				}
				err = bucket.PutObjectFromFile(core.DefaultConfig.App.Oss.Path+"test.xlsx", "./docs/test.xlsx")
				if err != nil {
					fmt.Println("the err0 is: ", err)
					return
				}
				return
			case <-c.Done():
				log.L.ErrorCtx(ctx, c.Err())
				err = c.Err()
				return
			default:
			}
		}
	}()
	return
}

//
//  GetReport
//  @Description: 从oss下载报表
//  @receiver s
//  @param ctx
//  @return code
//  @return err
//
func (r *ReportService) GetReport(ctx *trace.Trace) (code int, err error) {
	bucketName := core.DefaultConfig.App.Oss.BucketName
	bucket, err := oss.DefaultOssClient.Bucket(bucketName)
	if err != nil {
		return
	}
	go func(bucket *oss2.Bucket) {
		err = bucket.GetObjectToFile(core.DefaultConfig.App.Oss.Path+"test.xlsx", "./docs/oss-test.xlsx")
		if err != nil {
			return
		}
	}(bucket)
	return
}

//
//  GenerateReportByKafka
//  @Description: 接收kafka数据生成报表并上传
//  @receiver s
//  @param c
//  @return code
//  @return err
//
func (s *ReportService) GenerateReportByKafka(ctx *trace.Trace, req report.StartFlag) (code int, err error) {
	var (
		ch        = make(chan struct{}, 1)
		kafkaCh   = make(chan struct{}, 1)
		res, body []byte
		t         report.KafkaRes
		wait      sync.WaitGroup
	)
	defer close(ch)
	defer close(kafkaCh)
	data := make([]*report.Company, 0)
	c, cancelFunc := context.WithTimeout(ctx.Context, 5*time.Second)
	defer cancelFunc()
	//通知下游服务写数据到kafka中
	go func() {
		body, err = util.DoHttp(ch, _const.Post, _const.UrlKafkaData, req)
		if err != nil {
			log.L.ErrorCtx(ctx, err)
			return
		}
	}()
	wait.Add(1)
	go func() {
		defer wait.Done()
	loop:
		for {
			select {
			case <-ch:
				if string(body) == "" {
					err = errors.New("返回body为空，请检查")
					return
				}
				err = json.Unmarshal(body, &t)
				if err != nil {
					log.L.ErrorCtx(ctx, err)
					break loop
				}
				if t.Data.Status != 1 {
					err = errors.New("返回状态错误，请检查！")
					log.L.ErrorCtx(ctx, err)
					break loop
				}
				kafkaCtx, cancelFunc1 := context.WithTimeout(ctx.Context, 1200*time.Second)
				defer cancelFunc1()
				res, err = kafka.GetConsumerData(kafkaCh, core.DefaultConfig.App.Kafka.Address+":"+core.DefaultConfig.App.Kafka.Port, "test3")
				if err != nil {
					log.L.ErrorCtx(ctx, err)
					break loop
				}
				for {
					select {
					case <-kafkaCh:
						if len(res) == 0 {
							err = errors.New("kafka获取数据为空，请检查！")
							break loop
						}
						err = json.Unmarshal(res, &data)
						if err != nil {
							log.L.ErrorCtx(ctx, err)
							break loop
						}
						file := excelize.NewFile()
						defer file.Close()
						file.NewSheet("kafka-data")
						file.DeleteSheet("Sheet1")
						rowsCount := len(data)
						for i := 1; i <= rowsCount; i++ {
							_ = file.SetCellValue("kafka-data", fmt.Sprintf("A%d", i), data[i-1].Name)
							_ = file.SetCellValue("kafka-data", fmt.Sprintf("B%d", i), data[i-1].Employee)
							_ = file.SetCellValue("kafka-data", fmt.Sprintf("C%d", i), data[i-1].Email)
							_ = file.SetCellValue("kafka-data", fmt.Sprintf("D%d", i), data[i-1].Remark)
							_ = file.SetCellValue("kafka-data", fmt.Sprintf("D%d", i), data[i-1].Salary)
						}
						bucketName := core.DefaultConfig.App.Oss.BucketName
						bucket, err1 := oss.DefaultOssClient.Bucket(bucketName)
						if err1 != nil {
							log.L.ErrorCtx(ctx, err)
							break loop
						}
						err = file.SaveAs("./docs/kafka_data.xlsx")
						if err != nil {
							log.L.ErrorCtx(ctx, err)
							break loop
						}
						err = bucket.PutObjectFromFile(core.DefaultConfig.App.Oss.Path+"kafka_data.xlsx", "./docs/kafka_data.xlsx")
						if err != nil {
							log.L.ErrorCtx(ctx, err)
							break loop
						}
						//删除文件
						os.Remove("./docs/kafka_data.xlsx")
						break loop
					case <-kafkaCtx.Done():
						log.L.ErrorCtx(ctx, err)
						err = c.Err()
						break loop
					default:
					}
				}
			case <-c.Done():
				log.L.ErrorCtx(ctx, err)
				err = c.Err()
				return
			default:
			}
		}
	}()
	wait.Wait()
	return
}

//
// SendToElastic
//  @Description: 发送数据至elastic
//  @receiver s
//  @param ctx
//  @param req
//  @return code
//  @return err
//
func (s *ReportService) SendToElastic(ctx *trace.Trace, req report.ElasticReq) (code int, err error) {
	client := elastic.DefaultElasticClient
	_, err = client.Index().
		Index("user").
		BodyJson(req).
		Do(ctx)
	if err != nil {
		return
	}
	return
}
