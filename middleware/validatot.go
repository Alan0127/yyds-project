package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
	"yyds-pro/core"
	"yyds-pro/trace"
)

const (
	SUCCESS = 200
	ERROR   = 500

	//code = 1000 用户错误模块
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT    = 1008
)

var codeMsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户名已经存在",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_EXIST:      "token不存在，请重新登陆",
	ERROR_TOKEN_RUNTIME:    "token已过期，请重新登陆",
	ERROR_TOKEN_WRONG:      "token不正确，请重新登陆",
	ERROR_TOKEN_TYPE_WRONG: "token格式错误，请重新登陆",
	ERROR_USER_NO_RIGHT:    "该用户无权限",
}

var Jkey string
var JwtKey = []byte(Jkey)

type MyClaim struct {
	Username string `gorm:"username"`
	jwt.StandardClaims
}

//
//  SetToken
//  @Description: token生成函数
//  @param username
//  @return string
//  @return int
//
func SetToken(username string) (string, int) {
	expireTime := time.Now().Add(10 * time.Hour)
	SetCClaims := MyClaim{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "my_pro",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetCClaims)
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", ERROR
	}
	return token, SUCCESS

}

//
//  CheckToken
//  @Description: token验证
//  @param token
//  @return *MyClaim
//  @return int
//
func CheckToken(token string) (*MyClaim, int) {
	setToken, _ := jwt.ParseWithClaims(token, &MyClaim{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, code := setToken.Claims.(*MyClaim); code && setToken.Valid {
		return key, SUCCESS
	} else {
		return nil, ERROR
	}
}

//
//  JwtToken
//  @Description: jwt中间件
//  @return gin.HandlerFunc
//
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, traceCtx := core.GetTrace(c)
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			SetValidateTrace(ERROR_TOKEN_EXIST, traceCtx)
			c.Abort()
			return
		}
		checkToken := strings.Split(tokenHeader, " ")
		if len(checkToken) == 0 {
			SetValidateTrace(ERROR_TOKEN_TYPE_WRONG, traceCtx)
			c.Abort()
			return
		}
		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			SetValidateTrace(ERROR_TOKEN_TYPE_WRONG, traceCtx)
			c.Abort()
			return
		}
		key, Tcode := CheckToken(checkToken[1])
		if Tcode == ERROR {
			SetValidateTrace(ERROR_TOKEN_WRONG, traceCtx)
			c.Abort()
		}
		if time.Now().Unix() > key.ExpiresAt {
			SetValidateTrace(ERROR_TOKEN_RUNTIME, traceCtx)
			c.Abort()
		}
		c.Set("username", key.Username)
		c.Set(traceCtx.TraceId, traceCtx)
		c.Next()
	}
}

func GetErrorMsg(code int) string {
	return codeMsg[code]
}

func SetValidateTrace(code int, trace *trace.Trace) *trace.Trace {
	trace.Response.ErrorCode = code
	trace.Response.ErrorMessage = GetErrorMsg(code)
	return trace
}
