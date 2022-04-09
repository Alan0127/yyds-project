package repoimpl

import (
	"gorm.io/gorm"
	"yyds-pro/server/mysql"
	"yyds-pro/trace"
)

type ApkRepository struct {
	ApkDb *gorm.DB
}

func NewApkRepository() *ApkRepository {
	return &ApkRepository{
		ApkDb: mysql.Client,
	}
}

func (a ApkRepository) FindApkById(ctx *trace.Trace, id int) {
}
