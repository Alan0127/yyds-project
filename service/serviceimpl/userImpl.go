package serviceimpl

import (
	"yyds-pro/model"
	"yyds-pro/repository"
	"yyds-pro/repository/repoimpl"
	"yyds-pro/trace"
)

type UserService struct {
	UserRepo repository.UserRepo
}

func NewUserService() UserService {
	return UserService{
		UserRepo: repoimpl.NewUserRepository(),
	}
}

//
//  CheckLogin
//  @Description: 登录检查
//  @receiver s
//  @param ctx
//  @param user
//  @return model.UserInfo
//  @return error
//
func (s UserService) CheckLogin(ctx *trace.Trace, userName, userCountry string) (model.UserInfo, error) {
	return s.UserRepo.CheckLogin(ctx, userName, userCountry)
}

//
//  RegisterUser
//  @Description: 用户注册
//  @param ctx
//  @param RegisterUser
//  @return int
//  @return error
//
func (s UserService) RegisterUser(ctx *trace.Trace, registerUser model.RegisterInfo) (int, error) {
	return s.UserRepo.RegisterUser(ctx, registerUser)
}
