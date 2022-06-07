package repoimpl

import (
	"fmt"
	"gorm.io/gorm"
	"yyds-pro/log"
	"yyds-pro/server/mysql"
	"yyds-pro/trace"
)

type RateLimiterRepository struct {
	AppDb *gorm.DB
}

func NewRateLimiterRepository() *RateLimiterRepository {
	return &RateLimiterRepository{
		AppDb: mysql.Client,
	}
}

//
//  InitIntegralDb
//  @Description: 初始化积分信息表
//  @receiver r
//  @param ctx
//  @param id
//  @param money
//  @return err
//
func (r RateLimiterRepository) InitIntegralDb(ctx *trace.Trace, id, money int) (err error) {
	var i int
	err = r.AppDb.WithContext(ctx).Raw(`INSERT INTO integral_info
														(amount_id,
														 amount)
											VALUES      (?,
														 ?)`, id, money).Scan(&i).Error
	if err != nil {
		log.L.ErrorCtx(ctx, err)
	}
	return
}

//
//  UpdateUserIntegral
//  @Description: 更新user积分
//  @receiver r
//  @param ctx
//  @param val
//  @return err
//
func (r RateLimiterRepository) UpdateUserIntegral(ctx *trace.Trace, val int, userName string) (err error) {
	fmt.Println(userName)
	var i int
	err = r.AppDb.WithContext(ctx).Raw(`update 
											  user_info u 
											set 
											  u.user_integral = u.user_integral + ? where user_name = ?`, val, userName).Scan(&i).Error
	if err != nil {
		log.L.ErrorCtx(ctx, err)
	}
	return
}
