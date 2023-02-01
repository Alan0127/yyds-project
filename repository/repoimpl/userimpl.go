package repoimpl

import (
	"gorm.io/gorm"
	"yyds-pro/model"
	"yyds-pro/server/mysql"
	"yyds-pro/trace"
)

type UserRepository struct {
	AppDb *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		AppDb: mysql.Client,
	}
}

//
//  CheckLogin
//  @Description: 登录检查用户是否注册
//  @receiver u
//  @param ctx
//  @param user
//  @return res
//  @return err
//
func (u UserRepository) CheckLogin(ctx *trace.Trace, userName, userCountry string) (res model.UserResInfo, err error) {
	err = u.AppDb.WithContext(ctx).Raw(`SELECT * 
											FROM user_info 
											WHERE user_name= ?`, userName).Scan(&res).Error
	return
}

//
//  Register
//  @Description: 注册
//  @receiver u
//  @param ctx
//  @param user
//  @return res
//  @return err
//
func (u UserRepository) RegisterUser(ctx *trace.Trace, user model.RegisterInfo) (res int, err error) {
	err = u.AppDb.WithContext(ctx).Raw(`insert into user_info (
											  user_name, user_pass, user_country, 
											  user_age, user_gender, user_phone, 
											  user_wechat
											) 
											values 
											  (?, ?, ?, ?, ?, ?, ?)`, user.UserName,
		user.UserPass, user.Country, user.UserAge,
		user.UserGender, user.UserPhone, user.UserWechat).Scan(&res).Error
	return
}
