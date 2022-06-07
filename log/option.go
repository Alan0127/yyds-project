package log

import (
	"go.uber.org/zap"
	"path/filepath"
)

type Options struct {
	DevFlag      bool
	FileDir      string
	MaxSize      int //文件切分
	MaxBackUp    int //文件保留个数
	MaxAge       int
	Level        string
	CtxKey       string
	WriteFile    bool
	WriteConsole bool
}

type ZapLogOption func(option *Options)

func NewOption(opts ...ZapLogOption) *Options {
	opt := &Options{
		DevFlag:      true,
		MaxSize:      100,
		MaxBackUp:    60,
		MaxAge:       30,
		CtxKey:       "zapLog_key",
		WriteFile:    false,
		WriteConsole: true,
	}
	for _, f := range opts {
		f(opt)
	}
	return opt
}

func (l *Logger) LoadCfg() {
	if l.opts.DevFlag {
		l.zapConfig = zap.NewDevelopmentConfig()
		l.zapConfig.DisableStacktrace = true

	} else {
		l.zapConfig = zap.NewProductionConfig()
	}
}

func SetLevel(level string) ZapLogOption {
	return func(option *Options) {
		option.Level = level
	}
}

func SetPath(path string) ZapLogOption {
	logFileDir, _ := filepath.Abs(filepath.Dir(filepath.Join(".")))
	path = logFileDir + path
	return func(option *Options) {
		option.FileDir = path
	}
}
