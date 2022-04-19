package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
	"yyds-pro/log"
	trace2 "yyds-pro/trace"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		l := log.GetLogger()
		trace := trace2.NewTraceContext(c)
		start := time.Now()
		c.Set("traceId", trace.TraceId) //记录trace到Context中
		c.Set(trace.TraceId, trace)
		c.Next()
		requestUrl := c.Request.RequestURI
		method := c.Request.Method
		latency := time.Now().Sub(start).Seconds() //请求时间
		returnCode, _ := c.Get("returnCode")       //返回状态码
		//记录trace信息
		trace.WithReqUrl(requestUrl).
			WithMethod(method).
			WithLatency(latency).WithCode(returnCode.(int))
		l.InfoCtx(trace)
	}
}
