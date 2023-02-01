package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"yyds-pro/model"
)

var DefaultOssClient *oss.Client

//oss客户端初始化
func InitOss(config model.AppConfig) (err error) {
	endPoint := config.App.Oss.Endpoint
	accessKeyId := config.App.Oss.AccessKeyId
	accessKeySecret := config.App.Oss.AccessKeySecret
	DefaultOssClient, err = oss.New(endPoint, accessKeyId, accessKeySecret, oss.EnableCRC(false), oss.Timeout(10, 120))
	if err != nil {
		return
	}
	return
}
