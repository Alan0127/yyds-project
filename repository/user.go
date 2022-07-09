package repository

import (
	"yyds-pro/model"
	"yyds-pro/trace"
)

//
//  UserRepo
//  @Description: 登录repo interface
//
type UserRepo interface {
	//
	//  CheckLogin
	//  @Description: 用户检查
	//  @param ctx
	//  @param userName
	//  @param userCountry
	//  @return model.UserInfo
	//  @return error
	//
	CheckLogin(ctx *trace.Trace, userName, userCountry string) (model.UserInfo, error)

	//
	//  RegisterUser
	//  @Description: 用户注册
	//  @param ctx
	//  @param register
	//  @return int
	//  @return error
	//
	RegisterUser(ctx *trace.Trace, register model.RegisterInfo) (int, error)
}
