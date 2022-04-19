package repoimpl

import (
	"gorm.io/gorm"
	"yyds-pro/model"
	"yyds-pro/server/mysql"
	"yyds-pro/trace"
)

type AppRepository struct {
	AppDb *gorm.DB
}

func NewApkRepository() *AppRepository {
	return &AppRepository{
		AppDb: mysql.Client,
	}
}

func (a AppRepository) FindApkById(ctx *trace.Trace, id int) (res model.App, err error) {
	err = a.AppDb.WithContext(ctx).Raw(`select 
												app_name,
												app_online_time,
												app_language,
												app_desc 
											from app_info 
											where app_id = ?`, id).Scan(&res).Error
	return
}
