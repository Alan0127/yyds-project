package log

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
	"yyds-pro/common"
	"yyds-pro/trace"
)

var (
	L        *Logger             //全局log实例
	outWrite zapcore.WriteSyncer // IO输出
)

type Logger struct {
	*zap.Logger
	opts      *Options
	zapConfig zap.Config
}

func InitDefaultLog(opt ...ZapLogOption) {
	L = &Logger{
		opts: NewOption(opt...),
	}
	L.LoadCfg() //默认设置
	L.initLog()
}

//
//  initLog
//  @Description: log的初始化
//  @receiver l
//
func (l *Logger) initLog() {
	outWrite = zapcore.AddSync(&lumberjack.Logger{
		Filename:   l.opts.FileDir + "clean-log-info.log",
		MaxSize:    l.opts.MaxSize,
		MaxBackups: l.opts.MaxBackUp,
		MaxAge:     l.opts.MaxAge,
		Compress:   true,
		LocalTime:  true,
	})
	var err error
	l.Logger, err = l.zapConfig.Build(l.cores()) //返回一个log实例
	if err != nil {
		panic(err)
	}
	defer l.Logger.Sync()
}

func (l *Logger) cores() zap.Option {
	encoder := zapcore.NewJSONEncoder(l.zapConfig.EncoderConfig)    //使用json格式的编码器
	priority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool { //日志级别
		return lvl >= l.GetLevel()
	})
	var cores []zapcore.Core
	cores = append(cores, zapcore.NewCore(encoder, outWrite, priority))
	//WrapCore包装或替换Logger的底层zapcore.Core
	return zap.WrapCore(func(c zapcore.Core) zapcore.Core { //返回的是core列表
		return zapcore.NewTee(cores...)
	})
}

func (l *Logger) GetLevel() (level zapcore.Level) {
	switch strings.ToLower(l.opts.Level) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel //默认为调试模式
	}
}

func GetLogger() *Logger {
	if L == nil {
		fmt.Println("Please initialize the hlog service first")
		return nil
	}
	return L
}

//
//  InfoCtx
//  @Description: trace打印info日志
//  @receiver l
//  @param ctx
//
func (l *Logger) InfoCtx(ctx *trace.Trace) {
	fields := make([]zap.Field, 0)
	fields = append(fields,
		zap.Any("flag", ctx.Flag),
		zap.Any("traceId", ctx.TraceId),
		zap.Any("request", ctx.Req),
		zap.Any("response", ctx.Response),
		zap.Any("redis", ctx.Redis),
		zap.Any("sql", ctx.Sql),
	)
	l.Info("trace ", fields...)
}

//
//  ErrorCtx
//  @Description: trace打印Error日志
//  @receiver l
//  @param ctx
//  @param err
//
func (l *Logger) ErrorCtx(ctx *trace.Trace, err error) {
	file, line, pcName := common.Caller(2)
	ctx.Stack = fmt.Sprintf("%s, %s, %d", pcName, file, line)
	fields := make([]zap.Field, 0)
	fields = append(fields,
		zap.Any("flag", ctx.Flag),
		zap.Any("error", err),
		zap.Any("Stack", ctx.Stack),
	)
	l.Error("trace", fields...)
}
