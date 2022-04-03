package log

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

var (
	l        Logger
	outWrite zapcore.WriteSyncer // IO输出
)

type Logger struct {
	*zap.Logger
	opts      *Options
	zapConfig zap.Config
}

func InitDefaultLog(opt ...ZapLogOption) {
	l := &Logger{
		opts: NewOption(opt...),
	}
	fmt.Println(l)
	l.initLog()
}

func (l *Logger) initLog() {
	defer l.Logger.Sync()
	outWrite = zapcore.AddSync(&lumberjack.Logger{
		Filename:   l.opts.FileDir + "/" + "clean-arch" + ".log",
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

}

func (l *Logger) cores() zap.Option {
	encoder := zapcore.NewJSONEncoder(l.zapConfig.EncoderConfig)    //使用json格式的编码器
	priority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool { //日志级别
		return lvl >= l.GetLevel()
	})
	var cores []zapcore.Core
	cores = append(cores, []zapcore.Core{
		zapcore.NewCore(encoder, outWrite, priority), //ioCore
	}...)

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
