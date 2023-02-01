package _const

const (
	ResponseSuccess = 0
	ResponseError   = 1
	InitId          = 202108
)

var LanguageMap = make(map[string]bool)

const (
	TraceId              = "trace_id"
	GetAllAppsMsg        = "GetAllApps success!"
	GetApkByIdMsg        = "GetApkById success!"
	ChangeOrderStatusMsg = "ChangeOrderStatus success!"
	UserPurchaseMsg      = "UserPurchase success!"
	InitIntegralDbMsg    = "InitIntegralDb success!"
	SendKafkaSuccessful  = "sendKafkaSuccessful"
)

const (
	NIntegralList   = "N_INTEGRAL_LIST"
	YIntegralList   = "Y_INTEGRAL_LIST"
	IntegralInfo    = "INTEGRAL_INFO"
	AppInfosFiled   = "APP_INFOS_FIELD"
	AppInfos        = "APP_INFOS_%s"
	Key1            = "filter"
	UserKey         = "hasUser"
	HeaderToken     = "login_token"
	RedisLoginToken = "login-user"
	SessionUserInfo = "sessionUserInfo"
)

const (
	Post   = "POST"
	Get    = "GET"
	Delete = "DELETE"
	Put    = "PUT"
)

const (
	UrlPost      = "http://127.0.0.1:8800/httpService/getData"
	UrlKafkaData = "http://127.0.0.1:8800/httpService/getKafkaData"
)

func InitDefaultLanguage() {
	LanguageMap["en"] = true
	LanguageMap["zh_cn"] = true
}
