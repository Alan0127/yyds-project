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

//
//  FindApkById
//  @Description: FindApkById具体实现，根据id查询数据返回
//  @receiver a
//  @param ctx
//  @param id
//  @return res
//  @return err
//
func (a AppRepository) FindApkById(ctx *trace.Trace, id int) (res model.AppInfo, err error) {
	err = a.AppDb.WithContext(ctx).Raw(`select 
												id,
												app_name,
												app_status,
												app_type,
												app_version,
												app_img_id,
												app_video_id,
												app_link ,
												app_online_time,
												app_update_time
											from app_info 
											where id = ?`, id).Scan(&res).Error
	return
}

func (a AppRepository) ChangeTaskOrderStatusByOrderInfo(ctx *trace.Trace, orderReq model.OrderReq, cal uint) (err error) {
	err = a.AppDb.WithContext(ctx).Raw(`UPDATE 
											  task_order_user t 
											SET 
											  t.order_num = ? 
											WHERE 
											  t.user_id = ? 
											  AND t.task_id = ?`, cal, orderReq.UerId, orderReq.TaskId).Error
	return
}
