package repoimpl

import (
	"gorm.io/gorm"
	"yyds-pro/server/mysql"
)

type ReportRepository struct {
	AppDb *gorm.DB
}

func NewReportRepository() *ReportRepository {
	return &ReportRepository{
		AppDb: mysql.Client,
	}
}
