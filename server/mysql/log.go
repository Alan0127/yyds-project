package mysql

import (
	"context"
	"go.uber.org/zap"
	gormLogger "gorm.io/gorm/logger"
	"strings"
	"time"
	"yyds-pro/log"
	"yyds-pro/trace"
)

type Logger struct {
	ZapLogger                 *zap.Logger
	LogLevel                  gormLogger.LogLevel
	SlowThreshold             time.Duration
	SkipCallerLookup          bool
	IgnoreRecordNotFoundError bool
}

func New() Logger {
	return Logger{
		ZapLogger:                 log.GetLogger().Logger,
		LogLevel:                  gormLogger.Info,
		SlowThreshold:             100 * time.Millisecond,
		SkipCallerLookup:          false,
		IgnoreRecordNotFoundError: false,
	}
}

func (l Logger) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	return Logger{
		ZapLogger:                 l.ZapLogger,
		SlowThreshold:             l.SlowThreshold,
		LogLevel:                  level,
		SkipCallerLookup:          l.SkipCallerLookup,
		IgnoreRecordNotFoundError: l.IgnoreRecordNotFoundError,
	}
}

func (l Logger) Info(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormLogger.Info {
		return
	}
	log.L.Sugar().Debugf(str, args...)
}

func (l Logger) Warn(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormLogger.Warn {
		return
	}
	log.L.Sugar().Debugf(str, args...)
}

func (l Logger) Error(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormLogger.Error {
		return
	}
	log.L.Sugar().Debugf(str, args...)
}

//
//  Trace
//  @Description: 日志追踪
//  @receiver l
//  @param ctx
//  @param begin
//  @param fc
//  @param err
//
func (l Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	traceCtx, ok := ctx.(*trace.Trace)
	if !ok {
		log.L.Error("ctx error!", zap.Any("ctx", traceCtx))
	}
	sql, _ := fc()
	str := strings.Replace(strings.Replace(sql, "\n", "", -1), "\t", "", -1)
	if err != nil {
		traceCtx.Sql.Err = err
		traceCtx.Sql.SqlStr = str
		return
	}
	if l.LogLevel <= 0 {
		return
	}
	elapsed := time.Since(begin)
	traceCtx.Sql.SqlStr = str
	traceCtx.Sql.SqlElapsedTime = elapsed.Seconds()

}
