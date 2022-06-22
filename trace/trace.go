package trace

import (
	"context"
	"github.com/gin-gonic/gin"
	"yyds-pro/common"
	_const "yyds-pro/core/const"
)

const traceIDHeader = "X-TRACE-ID"

type Trace struct {
	context.Context
	TraceId  string
	Sql      Sql
	Redis    RedisRes
	Req      Request
	Response Response
	Flag     bool //请求是否成功
	Stack    string
}

type Sql struct {
	SqlStr         string  `json:"sql"`
	SqlElapsedTime float64 `json:"sqlElapsedTime"`
	Err            error   `json:"error"`
}

// Request 请求信息
type Request struct {
	ReqUrl string      `json:"reqUrl"`
	Method string      `json:"method"` // 请求方式
	Body   interface{} `json:"body"`   // 请求参数
}

type RedisRes struct {
	Error interface{} `json:"error"`
	Flag  bool        `json:"flag"`
}

// Response 响应信息
type Response struct {
	ErrorCode    interface{} `json:"errorCode"`    // 提示信息
	ErrorMessage string      `json:"errorMessage"` // HTTP 状态码
	Data         interface{} `json:"data"`         // HTTP 状态码信息
	CostSeconds  float64     `json:"costSecond"`   // 执行时间(单位秒)
}

func (t *Trace) WithResErrorMessage(err error) *Trace {
	t.Response.ErrorMessage = err.Error()
	return t
}

func (t *Trace) WithResErrorCode() *Trace {
	t.Response.ErrorCode = _const.ResponseError
	return t
}

func (t *Trace) WithResErrorFlag() *Trace {
	t.Flag = false
	return t
}

func (t *Trace) WithSuccessCode() *Trace {
	t.Response.ErrorCode = _const.ResponseSuccess
	return t
}

func (t *Trace) WithSuccessData(data interface{}) *Trace {
	t.Response.Data = data
	return t
}

func (t *Trace) WithSuccessFlag() *Trace {
	t.Flag = true
	return t
}

func NewTraceContext(ctx *gin.Context) *Trace {
	//查traceId
	traceID := ctx.Request.Header.Get(traceIDHeader)
	if traceID == "" {
		traceID = common.RandStringRunes(16)
	}
	trace := &Trace{
		Context: ctx,
		TraceId: traceID,
	}
	return trace
}

func (tc *Trace) WithReqUrl(url string) *Trace {
	tc.Req.ReqUrl = url
	return tc
}

func (tc *Trace) WithMethod(method string) *Trace {
	tc.Req.Method = method
	return tc
}

func (tc *Trace) WithLatency(latency float64) *Trace {
	tc.Response.CostSeconds = latency
	return tc
}

func (tc *Trace) WithTraceId(traceId string) *Trace {
	tc.TraceId = traceId
	return tc
}

func WithErrTrace(trace *Trace, err error) {
	trace.Response.ErrorCode = _const.ResponseSuccess
	if err != nil {
		trace.Flag = false
		trace.Response.ErrorMessage = err.Error()
	}
}
