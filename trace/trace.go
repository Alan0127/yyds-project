package trace

import "context"

type TContext struct {
	context.Context
	TraceId    string
	ReqUrl     string
	ReturnCode int
}

func NewTraceContext(ctx context.Context) *TContext {

	//初始化数据库

	//初始化redis

	trace := &TContext{
		Context: ctx,
	}
	return trace
}
