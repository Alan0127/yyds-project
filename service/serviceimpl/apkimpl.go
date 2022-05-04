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

func (s ApkService) GetApkById(ctx *trace.Trace, id int) (model.AppInfo, error) {
	return s.ApkRepo.FindApkById(ctx, id)
}
