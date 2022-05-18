package serviceimpl

import (
	"yyds-pro/model"
	"yyds-pro/repository"
	"yyds-pro/repository/repoimpl"
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
	var cal uint
	cal = 1
	//取消预约
	if orderReq.OrderReqType == 0 {
		cal = -1
	}
	return s.ApkRepo.ChangeTaskOrderStatusByOrderInfo(ctx, orderReq, cal)
}
