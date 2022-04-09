package core

import (
	"errors"
	"github.com/gin-gonic/gin"
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
	traceId, _ := c.Get("traceId")
	ctx, _ := c.Get(traceId.(string))
	if traceCtx, ok = ctx.(*trace.Trace); ok {
		return
	} else {
		return errors.New("trace error"), traceCtx
	}
}
