package model

//
//  UserInfo
//  @Description: 登录信息
//
type UserInfo struct {
	UserName string `json:"userName" gorm:"column:user_name" binding:"required"`
	Password string `json:"paasWord" gorm:"column:user_pass" binding:"required"`
	Country  string `json:"country"  gorm:"column:user_country" binding:"required"`
	//Token    string `json:"token"`
}

//
//  RegisterInfo
//  @Description: 注册信息
//
type RegisterInfo struct {
	UserName   string `json:"userName"     binding:"required"`
	UserPass   string `json:"userPass"     binding:"required"`
	Country    string `json:"country"      binding:"required"`
	UserAge    int64  `json:"userAge"      binding:"required"`
	UserGender string `json:"userGender"   binding:"required"`
	UserPhone  string `json:"userPhone"    binding:"required"`
	UserWechat string `json:"userWechat"   binding:"required"`
}
