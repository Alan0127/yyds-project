package service

import (
	"yyds-pro/model"
	"yyds-pro/trace"
)

//
//  UserService
//  @Description: 登录相关service interface
//
type UserService interface {
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
	//  @param user
	//  @return int
	//  @return error
	//
	RegisterUser(ctx *trace.Trace, user model.RegisterInfo) (int, error)
}
