package core

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	_const "yyds-pro/core/const"
	"yyds-pro/trace"
)

//
//  GetTrace
//  @Description: 从context中获得trace
//  @param c
//  @return err
//  @return traceCtx
//
func GetTrace(c *gin.Context) (err error, traceCtx *trace.Trace) {
	var ok bool
	traceId, _ := c.Get(_const.TraceId)
	if len(traceId.(string)) == 0 {
		err = errors.New("get traceId error")
		return
	}
	ctx, _ := c.Get(traceId.(string))
	if traceCtx, ok = ctx.(*trace.Trace); ok {
		return
	} else {
		err = errors.New("trace error")
		return
	}
}

//
//  BindReqWithContext
//  @Description: 绑定参数
//  @param traceCtx
//  @param c
//  @param data
//  @return err
//
func BindReqWithContext(traceCtx *trace.Trace, c *gin.Context, data interface{}) (err error) {
	err = c.ShouldBindWith(data, binding.JSON)
	if err != nil {
		return
	}
	traceCtx.Req.Body = data
	return
}
