package serviceimpl

import (
	"yyds-pro/repository"
	"yyds-pro/repository/repoimpl"
)

type ApkService struct {
	ApkRepo repository.ApkRepo
}

func NewApkService() ApkService {
	return ApkService{
		ApkRepo: repoimpl.NewApkRepository(),
	}
}

func (s ApkService) GetApkById(id int) {
	s.ApkRepo.FindApkById(id)
}
