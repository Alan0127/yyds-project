package repoimpl

import (
	"gorm.io/gorm"
	"yyds-pro/log"
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

func (a AppRepository) GetAllApps(ctx *trace.Trace, req model.GetAppsReq) (res []model.AppInfos, err error) {
	err = a.AppDb.WithContext(ctx).Raw(`SELECT ai.app_type, ai.app_status, ai.app_version, ai.app_img_id, ai.app_video_id
												, ad.app_description, ad.app_name_lang
											FROM app_info ai
												LEFT JOIN app_desc ad ON ai.id = ad.app_id
											WHERE ad.app_language = ?`, req.Language).Scan(&res).Error
	if err != nil {
		log.L.ErrorCtx(ctx, err)
	}
	return
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
	if err != nil {
		log.L.ErrorCtx(ctx, err)
	}
	return
}

func (a AppRepository) ChangeTaskOrderStatusByOrderInfo(ctx *trace.Trace, orderReq model.OrderReq, cal int) (err error) {
	var i int
	err = a.AppDb.WithContext(ctx).Raw(`UPDATE 
											  task_order_user t 
											SET 
											  t.task_order_status = ?
											WHERE 
											  t.user_id = ? 
											  AND t.task_id = ?`, cal, orderReq.UerId, orderReq.TaskId).Scan(&i).Error
	if err != nil {
		log.L.ErrorCtx(ctx, err)
	}
	return
}

func (a AppRepository) GetTaskUserOrderStatus(ctx *trace.Trace, orderReq model.OrderReq) (status int, err error) {
	err = a.AppDb.WithContext(ctx).Raw(`select 
											  status 
											from 
											  task_order_user t 
											where 
											  t.task_id = ? 
											  and t.user_id = ?`, orderReq.TaskId, orderReq.UerId).Scan(&status).Error
	if err != nil {
		log.L.ErrorCtx(ctx, err)
	}
	return
}
