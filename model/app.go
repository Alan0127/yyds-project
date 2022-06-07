package model

type language string

type GetAppsReq struct {
	Language language `json:"language"`
}

type ReqId struct {
	Id int `json:"id"`
}

type OrderReq struct {
	TaskId       int `json:"taskId"`
	UerId        int `json:"userId"`
	OrderReqType int `json:"orderReqType"`
}

type OrderRes struct {
	TaskId          int    `json:"taskId"`
	UserId          int    `json:"userId"`
	OrderNum        int    `json:"orderNum"`
	TaskOrderStatus int    `json:"taskOrderStatus"`
	TaskDesc        string `json:"taskDesc"`
}

type AppInfo struct {
	Id            int    `json:"id" gorm:"column:id"`
	AppName       string `json:"appName" gorm:"column:app_name"`
	AppType       int    `json:"appType" gorm:"column:app_type"`
	AppStatus     int    `json:"appStatus" gorm:"column:app_status"`
	AppVersion    string `json:"appVersion" gorm:"column:app_version"`
	AppImgId      int    `json:"appImgId" gorm:"column:app_img_id"`
	AppVideoId    int    `json:"appVideoId" gorm:"column:app_video_id"`
	AppLink       int    `json:"appLink" gorm:"column:app_link"`
	AppOnlineTime string `json:"appOnlineTime" gorm:"column:app_online_time"`
	AppUpdateTime string `json:"appUpdateTime" gorm:"column:app_update_time"`
}

type AppInfos struct {
	AppType     int    `json:"appType" gorm:"column:app_type"`
	AppStatus   int    `json:"appStatus" gorm:"column:app_status"`
	AppVersion  string `json:"appVersion" gorm:"column:app_version"`
	AppImgId    int    `json:"appImgId" gorm:"column:app_img_id"`
	AppVideoId  int    `json:"appVideoId" gorm:"column:app_video_id"`
	AppDesc     string `json:"appDesc" gorm:"column:app_description"`
	AppNameLang string `json:"appNameLang" gorm:"column:app_name_lang"`
}

type UserPurchase struct {
	UserName      string `json:"userName" gorm:"column:user_name"`
	UserPass      string `json:"userPass" gorm:"column:user_pass"`
	UserCountry   string `json:"userCountry" gorm:"column:user_country"`
	ActivityId    int    `json:"taskId" column:"id"`
	TotalIntegral int    `json:"totalIntegral" gorm:"column:total_integral"`
}
