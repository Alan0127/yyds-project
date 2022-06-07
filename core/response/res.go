package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yyds-pro/core/const"
	"yyds-pro/trace"
)

//
//  ResSuccess
//  @Description: 正确返回
//  @param c
//  @param trace
//  @param data
//
func ResSuccess(c *gin.Context, trace *trace.Trace, data interface{}, msg string) {
	trace.Response.Data = data
	trace.Response.ErrorCode = _const.ResponseSuccess
	trace.Flag = true
	c.JSON(http.StatusOK, gin.H{
		"errorCode":    _const.ResponseSuccess,
		"errorMessage": msg,
		"data":         data,
	})
}

//
//  ResError
//  @Description: 错误返回
//  @param c
//  @param trace
//  @param err
//
func ResError(c *gin.Context, trace *trace.Trace, err error) {
	trace.Response.ErrorCode = _const.ResponseSuccess
	trace.Response.ErrorMessage = err.Error()
	trace.Flag = false
	c.JSON(http.StatusOK, gin.H{
		"errorCode":    _const.ResponseError,
		"errorMessage": err,
		"data":         "",
	})
}
