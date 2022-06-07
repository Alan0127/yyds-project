package serviceimpl

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"yyds-pro/core/const"
	"yyds-pro/log"
	"yyds-pro/model"
	"yyds-pro/repository"
	"yyds-pro/repository/repoimpl"
	defaultRedis "yyds-pro/server/redis"
	"yyds-pro/trace"
)

type ApkService struct {
	ApkRepo repository.ApkRepo
}

func NewApkService() ApkService {
	return ApkService{
		ApkRepo: repoimpl.NewApkRepository(),
	}
}

//
//  GetAllApps
//  @Description: 获取所有app详情信息返回
//  @receiver s
//  @param ctx
//  @param req
//  @return res
//  @return err
//
func (s ApkService) GetAllApps(ctx *trace.Trace, req model.GetAppsReq) (res []model.AppInfos, err error) {
	//先走缓存，没有缓存再走数据库查询，并更新缓存
	field := _const.AppInfosFiled
	key := fmt.Sprintf(_const.AppInfos, req.Language)
	v, err := defaultRedis.HashGetWithCtx(ctx, field, key)
	if err != nil && err != redis.Nil {
		log.L.ErrorCtx(ctx, err)
		return
	}
	if len(v) != 0 {
		err = json.Unmarshal([]byte(v), &res)
		if err != nil {
			log.L.ErrorCtx(ctx, err)
			return
		}
		return
	}
	res, err = s.ApkRepo.GetAllApps(ctx, req)
	temp, _ := json.Marshal(&res)
	err = defaultRedis.HashSetWithContext(ctx, field, key, temp)
	if err != nil {
		log.L.ErrorCtx(ctx, err)
		return
	}
	return
}

//
//  GetApkById
//  @Description: GetApkById实现方法
//  @receiver s
//  @param ctx
//  @param id
//  @return model.AppInfo
//  @return error
//
func (s ApkService) GetApkById(ctx *trace.Trace, id int) (model.AppInfo, error) {
	return s.ApkRepo.FindApkById(ctx, id)
}

//
//  ChangeTaskOrderStatusByOrderInfo
//  @Description: ChangeTaskOrderStatusByOrderInfo实现方法
//  @receiver s
//  @param ctx
//  @param orderReq
//  @return model.OrderRes
//  @return error
//
func (s ApkService) ChangeTaskOrderStatusByOrderInfo(ctx *trace.Trace, orderReq model.OrderReq) (err error) {
	var cal int
	cal = 1
	//取消预约
	if orderReq.OrderReqType == 0 {
		cal = 0
	}
	status, err := s.ApkRepo.GetTaskUserOrderStatus(ctx, orderReq)
	if err != nil {
		return
	}
	if status == cal {
		return
	}
	return s.ApkRepo.ChangeTaskOrderStatusByOrderInfo(ctx, orderReq, cal)
}
