package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yyds-pro/core/code"
	"yyds-pro/trace"
)

//
//  ResSuccess
//  @Description: 正确返回
//  @param c
//  @param trace
//  @param data
//
func ResSuccess(c *gin.Context, trace *trace.Trace, data interface{}) {
	trace.Response.Data = data
	trace.Response.ErrorCode = code.ResponseSuccess
	trace.Flag = true
	c.JSON(http.StatusOK, gin.H{
		"errorCode":    code.ResponseSuccess,
		"errorMessage": "",
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
	trace.Response.ErrorCode = code.ResponseSuccess
	trace.Response.ErrorMessage = err.Error()
	trace.Flag = false
	c.JSON(http.StatusOK, gin.H{
		"errorCode":    code.ResponseError,
		"errorMessage": err,
		"data":         "",
	})
}
