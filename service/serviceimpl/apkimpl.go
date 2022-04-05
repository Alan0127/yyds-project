package serviceimpl

import (
	"yyds-pro/repository"
	"yyds-pro/repository/repoimpl"
)

var Apks ApkService

type ApkService struct {
	ApkRepo repository.ApkRepo
}

func init() {
	Apks = NewApkService()
}

func NewApkService() ApkService {
	return ApkService{
		ApkRepo: repoimpl.NewApkRepository(),
	}
}

func (s ApkService) GetApkById(id int) {
	s.ApkRepo.FindApkById(id)
}
